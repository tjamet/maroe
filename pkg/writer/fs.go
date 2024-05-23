package writer

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/tjamet/maroe/pkg/picture"
)

type FSWriter struct {
	fs afero.Fs
}

var _ Writer = &FSWriter{}

func NewFSWriter(fs afero.Fs) *FSWriter {
	return &FSWriter{fs: fs}
}

// Write a photo in the correct destination. Creates the directories when needed
func (w *FSWriter) Write(photo *picture.Picture, destination string) (string, error) {
	outPath := destPath(photo, destination)
	tmpPath := path.Join(filepath.Dir(outPath), "."+filepath.Base(outPath))
	err := w.fs.MkdirAll(filepath.Dir(outPath), 0700)
	if err != nil {
		return outPath, err
	}
	_, err = w.fs.Stat(outPath)
	if !os.IsNotExist(err) {
		return outPath, fmt.Errorf("path %s already exist", outPath)
	}
	fd, err := w.fs.Create(tmpPath)
	if err != nil {
		return outPath, fmt.Errorf("failed to create file %s: %v", tmpPath, err)
	}
	_, err = io.Copy(fd, photo.Reader)
	if err != nil {
		w.fs.Remove(tmpPath)
		return outPath, fmt.Errorf("failed to write file %s: %v", outPath, err)
	}
	err = w.fs.Rename(tmpPath, outPath)
	if err != nil {
		w.fs.Remove(tmpPath)
		return outPath, fmt.Errorf("failed to write file %s: %v", outPath, err)
	}
	return outPath, nil
}
