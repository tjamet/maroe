package writer

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/tjamet/maroe/pkg/picture"
)

// Write a photo in the correct destination. Creates the directories when needed
func Write(photo *picture.Picture, destination string) (string, error) {
	outPath := path.Join(destination, photo.Name+photo.Extension)
	tmpPath := path.Join(filepath.Dir(outPath), "."+filepath.Base(outPath))
	err := os.MkdirAll(filepath.Dir(outPath), 0700)
	if err != nil {
		return outPath, nil
	}
	_, err = os.Stat(outPath)
	if !os.IsNotExist(err) {
		return outPath, fmt.Errorf("path %s already exist", outPath)
	}
	fd, err := os.Create(tmpPath)
	if err != nil {
		return outPath, fmt.Errorf("failed to create file %s: %v", tmpPath, err)
	}
	_, err = io.Copy(fd, photo.Reader)
	if err != nil {
		os.Remove(tmpPath)
		return outPath, fmt.Errorf("failed to write file %s: %v", outPath, err)
	}
	err = os.Rename(tmpPath, outPath)
	if err != nil {
		os.Remove(tmpPath)
		return outPath, fmt.Errorf("failed to write file %s: %v", outPath, err)
	}
	return outPath, nil
}
