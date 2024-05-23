package lister

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tjamet/maroe/pkg/picture"
)

var allowedExtensions = []string{".cr2", ".jpg", ".jpeg", ".raf", ".nef"}

// Lister holds the context to list all photos in a path
type Lister struct {
	duplicateExtensionChecksums map[string]string
}

// Start listing photos matching supported extensions in the path
func (l *Lister) Start(path string) chan *picture.Picture {
	r := make(chan *picture.Picture, 2)
	go func() {
		defer close(r)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			log.Fatalf("could not walk non existing directory %s", path)
		}

		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return fmt.Errorf("Invalid file %s", path)
			}
			if info.IsDir() {
				return nil
			}
			ext := filepath.Ext(path)
			extLower := strings.ToLower(ext)
			for _, allowedExt := range allowedExtensions {
				if extLower == allowedExt {
					p, err := picture.NewPictureFromFilePath(path)
					if err != nil {
						return fmt.Errorf("failed to load picture %s: %s", path, err)
					}
					r <- p
					break
				}
			}
			return nil
		})
		if err != nil {
			log.Fatalf("Failed to find all photos: %s", err)
		}
	}()
	return r
}
