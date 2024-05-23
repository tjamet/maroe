package writer

import (
	"path"

	"github.com/tjamet/maroe/pkg/picture"
)

type Writer interface {
	Write(photo *picture.Picture, destination string) (string, error)
}

func destPath(photo *picture.Picture, destination string) string {
	return path.Join(destination, photo.Name+photo.Extension)
}
