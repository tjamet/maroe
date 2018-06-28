package picture

import (
	"io"
	"time"
)

type Picture struct {
	Name        string
	Extension   string
	CaptureTime time.Time
	Reader      io.Reader
	Basename    string
	Checksum    string
	original    string
}
