package archive

import (
	"go.uber.org/zap"
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
			Name: filepath.Base(srcFolder),
			Path: srcFolder,
		}

		return file, nil
	}

	zap.L().Debug(srcFolder + " is not a tar gzip file. Archiving and compressing now.")

	return ArchiveTGZ(srcFolder)
}