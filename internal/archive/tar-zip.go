package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

type TGZFile struct {
	Content []byte
	Name string
	Path string
}

func ArchiveTGZ(srcFolder string) (*TGZFile, error) {
	destFile := filepath.Clean(srcFolder) + ".tgz"
	tarGzFile, err := os.Create(destFile)
	if err != nil {
		return nil, err
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

	content, err := os.ReadFile(destFile)
	if err != nil {
		return nil, err
	}

	file := &TGZFile{
		Content: content,
		Name: filepath.Base(destFile),
		Path: destFile,
	}

	return file, nil
}