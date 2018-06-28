package copy

import (
	"log"
	"strings"

	"github.com/tjamet/maroe/pkg/writer"

	"github.com/tjamet/maroe/pkg/lister"
)

// Copy synchronizes source and destination using checksum names
func Copy(destination, source string) error {
	l := lister.Lister{}
	for photo := range l.Start(source) {
		path, err := writer.Write(photo, destination)
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
