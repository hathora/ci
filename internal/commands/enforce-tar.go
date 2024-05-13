package commands

import (
	"github.com/hathora/ci/internal/compress"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"slices"
)

var allowedExtensions = []string{"tgz", "gz"}

func isTarFar(filePath string) bool {
	ext := filepath.Ext(filePath)
	return slices.Contains(allowedExtensions, ext)
}

func EnforceTar(srcFolder string) (file []byte, err error) {
	if isTarFar(srcFolder) {
		zap.L().Info("File is already a tar file.")
		return os.ReadFile(srcFolder)
	}
	
	zap.L().Info(srcFolder)
	
	return compress.TarZip(srcFolder)
}