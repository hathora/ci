package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"

	"github.com/h2non/filetype"
	"github.com/monochromegane/go-gitignore"
	"go.uber.org/zap"
)

type TGZFile struct {
	Content []byte
	Name    string
	Path    string
}

func getIgnoreMatchers(srcFolder string, filepaths ...string) ([]gitignore.IgnoreMatcher, error) {
	var matchers []gitignore.IgnoreMatcher
	for _, path := range filepaths {
		matcher, err := gitignore.NewGitIgnore(filepath.Join(srcFolder, path), ".")
		if err != nil {
			zap.L().Debug("Could not find a " + path + " file. " + path + " matcher will not be used.")
			continue
		}

		matchers = append(matchers, matcher)
	}

	return matchers, nil
}

func shouldIgnoreFilepath(filepath string, isDir bool, matchers []gitignore.IgnoreMatcher) bool {
	anyMatches := false
	for _, matcher := range matchers {
		if matcher.Match(filepath, isDir) {
			anyMatches = true
			break
		}
	}

	return anyMatches
}

func ArchiveTGZ(srcFolder string) (string, error) {
	fileName := filepath.Base(filepath.Clean(srcFolder))
	destinationFile := fmt.Sprintf("%s.*.tgz", fileName)
	tarGzFile, err := os.CreateTemp("", destinationFile)
	if err != nil {
		return "", err
	}
	defer tarGzFile.Close()

	gzipWriter := gzip.NewWriter(tarGzFile)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	ignoreMatchers, err := getIgnoreMatchers(
		srcFolder,
		".dockerignore",
		".gitignore",
	)

	if err != nil {
		return "", err
	}

	err = filepath.Walk(srcFolder, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcFolder, filePath)
		if err != nil {
			return err
		}

		if shouldIgnoreFilepath(relPath, info.IsDir(), ignoreMatchers) {
			return nil
		}

		header, err := tar.FileInfoHeader(info, relPath)
		if err != nil {
			return err
		}

		header.Name = relPath

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := io.Copy(tarWriter, file); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return tarGzFile.Name(), nil
}

var (
	supportedFileExtensions = []string{".tgz", ".tar.gz", ".tar.tgz"}
)

func isTGZ(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}

	if fileInfo.IsDir() {
		return false, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	buff := make([]byte, 512)
	if _, err = file.Read(buff); err != nil {
		return false, fmt.Errorf("could not read file: %s", filePath)
	}

	if filetype.IsArchive(buff) {
		fileExt := filepath.Ext(filePath)
		if !slices.Contains(supportedFileExtensions, fileExt) {
			err := fmt.Errorf("unsupported archive file extension: %s", fileExt)
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func RequireTGZ(srcFolder string) (*TGZFile, error) {
	isFileTGZ, err := isTGZ(srcFolder)
	if err != nil {
		return nil, err
	}

	if isFileTGZ {
		zap.L().Debug(srcFolder + " is already a gzipped tar archive")
		content, err := os.ReadFile(srcFolder)
		if err != nil {
			return nil, err
		}
		file := &TGZFile{
			Content: content,
			Name:    filepath.Base(srcFolder),
			Path:    srcFolder,
		}

		return file, nil
	}

	zap.L().Debug(srcFolder + " is not a gzipped tar archive. Archiving and compressing now.")

	destFile, err := ArchiveTGZ(srcFolder)
	if err != nil {
		return nil, err
	}
	defer os.Remove(destFile)

	content, err := os.ReadFile(destFile)
	if err != nil {
		return nil, err
	}

	file := &TGZFile{
		Content: content,
		Name:    filepath.Base(destFile),
		Path:    destFile,
	}

	return file, nil
}
