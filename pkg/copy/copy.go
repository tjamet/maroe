package copy

import (
	"log"
	"strings"

	"github.com/tjamet/maroe/pkg/writer"

	"github.com/tjamet/maroe/pkg/lister"
)

type Copier interface {
	Copy(destination, source string) error
}

type copier struct {
	writer writer.Writer
}

var _ Copier = &copier{}

func NewCopier(w writer.Writer) Copier {
	return &copier{writer: w}
}

// Copy synchronizes source and destination using checksum names
func (c *copier) Copy(destination, source string) error {
	l := lister.Lister{}
	for photo := range l.Start(source) {
		path, err := c.writer.Write(photo, destination)
		if err != nil {
			if strings.Contains(err.Error(), "already exist") {
				log.Printf("Skipped copying %s to %s, destination already exist", photo.GetOriginalPath(), path)
			} else {
				log.Printf("Failed to copy %s to %s", photo.GetOriginalPath(), path)
			}
		} else {
			log.Printf("Copied %s to %s", photo.GetOriginalPath(), path)
		}
	}
	return nil
}
