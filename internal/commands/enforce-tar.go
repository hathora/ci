package commands

import (
	"github.com/hathora/ci/internal/compress"
	"net/http"
	"os"
	"path/filepath"
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
		return false, err
	}

	fileType := http.DetectContentType(buff)

	return fileType == "application/gzip", nil
}

func RequireTGZ(srcFolder string) (*compress.TGZFile, error) {
	isFileTGZ, err := isTGZ(srcFolder)
	if err != nil {
		return nil, err
	}

	if isFileTGZ {
		content, err := os.ReadFile(srcFolder)
		if err != nil {
			return nil, err
		}
		file := &compress.TGZFile{
			Content: content,
			Name: filepath.Base(srcFolder),
			Path: srcFolder,
		}

		return file, nil

	}

	return compress.ArchiveTGZ(srcFolder)
}
