package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

type TGZFile struct {
	Content []byte
	Name    string
	Path    string
}

func ArchiveTGZ(srcFolder string) (string, error) {
	destFile := filepath.Clean(srcFolder) + ".tgz"
	tarGzFile, err := os.Create(destFile)
	if err != nil {
		return "", err
	}
	defer tarGzFile.Close()

	gzipWriter := gzip.NewWriter(tarGzFile)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	err = filepath.Walk(srcFolder, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcFolder, filePath)
		if err != nil {
			return err
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

	return destFile, nil
}

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
		return false, err
	}

	fileType := http.DetectContentType(buff)

	return fileType == "application/gzip" || fileType == "application/x-gzip", nil
}

func RequireTGZ(srcFolder string) (*TGZFile, error) {
	isFileTGZ, err := isTGZ(srcFolder)
	if err != nil {
		return nil, err
	}

	if isFileTGZ {
		zap.L().Debug(srcFolder + " is already a tar gzip file")
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

	zap.L().Debug(srcFolder + " is not a tar gzip file. Archiving and compressing now.")

	destFile, err := ArchiveTGZ(srcFolder)
	if err != nil {
		return nil, err
	}

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
