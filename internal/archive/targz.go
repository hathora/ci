package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/boyter/gocodewalker"
	"github.com/h2non/filetype"
	"go.uber.org/zap"
)

type TGZFile struct {
	Content []byte
	Name    string
	Path    string
}

func CreateTGZ(srcFolder string, ext string) (string, error) {
	fileName := filepath.Base(filepath.Clean(srcFolder))
	destinationFile := fmt.Sprintf("%s.*.%s", fileName, ext)
	tarGzFile, err := os.CreateTemp("", destinationFile)
	if err != nil {
		return "", err
	}
	defer tarGzFile.Close()

	gzipWriter := gzip.NewWriter(tarGzFile)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	paths := make(chan *gocodewalker.File, 100)

	walker := gocodewalker.NewFileWalker(srcFolder, paths)
	// .gitignore is ignored by default
	walker.CustomIgnore = []string{".dockerignore"}
	walker.SetErrorHandler(func(err error) bool {
		zap.L().Error("Error walking directory", zap.Error(err))
		return false
	})

	err = walker.Start()
	if err != nil {
		return "", fmt.Errorf("error walking directory: %w", err)
	}

	for path := range paths {

		info, err := os.Stat(path.Location)
		if err != nil {
			return "", fmt.Errorf("error statting file: %w", err)
		}

		if !info.Mode().IsRegular() {
			continue
		}

		relPath, err := filepath.Rel(srcFolder, path.Location)
		if err != nil {
			return "", fmt.Errorf("error getting relative path: %w", err)
		}

		zap.L().Debug("Considering file: " + relPath)

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return "", fmt.Errorf("error creating tar file info header: %w", err)
		}
		header.Name = filepath.ToSlash(relPath)

		if err := tarWriter.WriteHeader(header); err != nil {
			return "", err
		}

		if info.IsDir() {
			zap.L().Debug("Including directory reference: " + relPath)
			continue
		}

		file, err := os.Open(path.Location)
		if err != nil {
			return "", fmt.Errorf("error opening file: %w", err)
		}
		defer file.Close()

		written, err := io.Copy(tarWriter, file)
		if err != nil {
			return "", fmt.Errorf("error copying file content for %s: %w", path.Location, err)
		}
		if written != info.Size() {
			return "", fmt.Errorf("expected to write %d bytes but wrote %d bytes for file %s", info.Size(), written, path.Location)
		}

		zap.L().Debug("Including file: " + relPath)
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

	if !filetype.IsArchive(buff) {
		return false, nil
	}

	var fileExtParts []string
	for {
		ext := filepath.Ext(filePath)
		if ext == "" {
			break
		}
		fileExtParts = append(fileExtParts, ext)
		filePath = filePath[:len(filePath)-len(ext)]
	}
	slices.Reverse(fileExtParts)

	var fileExt string
	for i := 0; i < len(fileExtParts); i++ {
		fileExt = strings.Join(fileExtParts[i:], "")
		if slices.Contains(supportedFileExtensions, fileExt) {
			return true, nil
		}
	}

	return false, fmt.Errorf("unsupported archive file extension: %s", fileExt)
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

	destFile, err := CreateTGZ(srcFolder, "tgz")
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
