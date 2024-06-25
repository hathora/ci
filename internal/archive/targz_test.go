package archive_test

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/hathora/ci/internal/archive"
)

func Test_CreateTGZ(t *testing.T) {
	zapLogger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(zapLogger)
	tests := []struct {
		name       string
		files      map[string]string
		shouldFail bool
	}{
		{
			name: "simple",
			files: map[string]string{
				"file1.txt":        "This is file 1",
				"subdir/file2.txt": "This is file 2 in subdir",
			},
			shouldFail: false,
		},
		{
			name: "nested",
			files: map[string]string{
				"dir1/dir2/file3.txt": "This is file 3 in nested directory",
				"dir1/file4.txt":      "This is file 4",
			},
			shouldFail: false,
		},
		{
			name: "nested with dirs only",
			files: map[string]string{
				"dir3/":               "",
				"dir3/dir4/":          "",
				"dir3/dir4/file5.txt": "This is file 5 in nested empty directory",
			},
			shouldFail: false,
		},
		{
			name: "special characters in filenames",
			files: map[string]string{
				"file with spaces.txt":  "This is a file with spaces",
				"file-with-üñîçødé.txt": "This file has unicode characters",
			},
			shouldFail: false,
		},
		{
			name:       "empty",
			files:      map[string]string{},
			shouldFail: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			srcFolder, err := os.MkdirTemp("", "testsrc")
			require.NoError(t, err)
			t.Cleanup(func() {
				os.RemoveAll(srcFolder)
			})

			for path, content := range tt.files {
				fullPath := filepath.Join(srcFolder, path)
				if strings.HasSuffix(path, "/") {
					require.NoError(t, os.MkdirAll(fullPath, os.ModePerm))
				} else {
					err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
					require.NoError(t, err)
					err = os.WriteFile(fullPath, []byte(content), 0644)
					require.NoError(t, err)
				}
			}

			archivePath, err := archive.CreateTGZ(srcFolder, "tgz")
			if tt.shouldFail {
				assert.Error(err)
				return
			} else {
				assert.NoError(err)
			}
			t.Cleanup(func() {
				os.Remove(archivePath)
			})

			file, err := os.Open(archivePath)
			assert.NoError(err)
			t.Cleanup(func() {
				file.Close()
			})

			gzipReader, err := gzip.NewReader(file)
			assert.NoError(err)
			t.Cleanup(func() {
				gzipReader.Close()
			})

			tarReader := tar.NewReader(gzipReader)

			archivedFiles := make(map[string]string)
			for {
				header, err := tarReader.Next()
				if err == io.EOF {
					break
				}
				assert.NoError(err)

				if header.Typeflag == tar.TypeReg {
					content, err := io.ReadAll(tarReader)
					assert.NoError(err)
					archivedFiles[header.Name] = string(content)
				}
			}

			for path, expectedContent := range tt.files {
				if strings.HasSuffix(path, "/") {
					continue // Skip directories
				}
				content, found := archivedFiles[path]
				assert.True(found, "Expected file %s not found in archive", path)
				assert.Equal(expectedContent, content, "File content mismatch for %s", path)
			}
		})
	}
}

func Test_RequireTGZ(t *testing.T) {
	zapLogger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(zapLogger)
	tests := []struct {
		name             string
		files            map[string]string
		archiveExt       string
		precreateArchive bool
	}{
		{
			name: "existing .tgz",
			files: map[string]string{
				"file1.txt": "This is file 1",
			},
			archiveExt:       "tgz",
			precreateArchive: true,
		},
		{
			name: "existing .tar.gz",
			files: map[string]string{
				"file1.txt": "This is file 1",
			},
			archiveExt:       "tar.gz",
			precreateArchive: true,
		},
		{
			name: "existing .tar.tgz",
			files: map[string]string{
				"file1.txt": "This is file 1",
			},
			archiveExt:       "tar.tgz",
			precreateArchive: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			srcFolder, err := os.MkdirTemp("", "testsrc")
			require.NoError(t, err)
			t.Cleanup(func() {
				os.RemoveAll(srcFolder)
			})

			for path, content := range tt.files {
				fullPath := filepath.Join(srcFolder, path)
				if strings.HasSuffix(path, "/") {
					require.NoError(t, os.MkdirAll(fullPath, os.ModePerm))
				} else {
					err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
					require.NoError(t, err)
					err = os.WriteFile(fullPath, []byte(content), 0644)
					require.NoError(t, err)
				}
			}

			var archivePath string
			if tt.precreateArchive {
				archivePath, err = archive.CreateTGZ(srcFolder, tt.archiveExt)
				require.NoError(t, err)
				t.Cleanup(func() {
					os.Remove(archivePath)
				})
			}

			tgzFile, err := archive.RequireTGZ(archivePath)
			require.NoError(t, err)
			require.NotNil(t, tgzFile)
			assert.Equal(archivePath, tgzFile.Path)
		})
	}
}
