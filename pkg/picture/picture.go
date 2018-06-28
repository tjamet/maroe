package picture

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tjamet/goraw"
	_ "github.com/tjamet/goraw/all" // enable all known file types
	"github.com/tjamet/maroe/image/tiff"
	"github.com/tjamet/maroe/image/tiff/exif"
)

// GetOriginalPath returns the path of the picture that has been loaded
func (p *Picture) GetOriginalPath() string {
	return p.original
}

func tryFindThumb(t *tiff.Tiff, offsetMarker, lengthMarker string) (start, size uint32, err error) {
	possInterface, err := t.DecodeFirst(offsetMarker)
	if err != nil {
		return 0, 0, err
	}
	poss, ok := possInterface.([]uint32)
	if !ok || len(poss) < 1 {
		return 0, 0, fmt.Errorf("failed to decode offset")
	}
	lengthsIntf, err := t.DecodeFirst(lengthMarker)
	if err != nil {
		return 0, 0, err
	}
	lengths, ok := lengthsIntf.([]uint32)
	if !ok || len(lengths) < 1 {
		return 0, 0, fmt.Errorf("failed to decode length")
	}
	if len(poss) != len(lengths) {
		return 0, 0, fmt.Errorf("offset and length do not have the same length")
	}
	smallest := 0
	// Find the smallest thumbnail
	for i := range lengths {
		if lengths[i] < lengths[smallest] {
			smallest = i
		}
	}
	smallest = 0
	return poss[smallest], lengths[smallest], nil
}

func findThumbnail(t *tiff.Tiff) (start, size uint32, err error) {
	start, length, err := tryFindThumb(t, exif.ThumbnailOffset, exif.ThumbnailLength)
	if err == nil {
		return start, length, nil
	}
	if !errors.Is(err, &tiff.ErrTagNotFound{}) {
		return 0, 0, err
	}
	start, length, err = tryFindThumb(t, exif.PreviewImageStart, exif.PreviewImageLength)
	if err == nil {
		return start, length, nil
	}
	if !errors.Is(err, &tiff.ErrTagNotFound{}) {
		return 0, 0, err
	}
	return 0, 0, errors.New("fid not find any thumbnail")
}

// NewPictureFromFilePath reads the whole file in memory to create a Picture handle
func NewPictureFromFilePath(path string) (*Picture, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	prefix := ""
	var captureTime time.Time
	h := sha1.New()
	ext := filepath.Ext(path)
	raw, err := goraw.New(fd)
	if err != nil {
		if err.Error() == "unsupported format" {
			fd.Seek(0, io.SeekStart)
			_, err := io.Copy(h, fd)
			if err != nil {
				return nil, fmt.Errorf("failed to get checksum for input file: %v", err)
			}
		} else {
			return nil, err
		}
	} else {
		exifReader, err := raw.ExifReaderAt()
		if err != nil {
			fd.Seek(0, io.SeekStart)
			_, err := io.Copy(h, fd)
			if err != nil {
				return nil, fmt.Errorf("failed to get checksum for input file: %v", err)
			}
		} else {
			t, err := tiff.Read(exifReader)
			if err != nil {
				return nil, err
			}
			thumbPos, thumbLen, err := findThumbnail(t)

			if err != nil {
				fd.Seek(0, io.SeekStart)
				_, err := io.Copy(h, fd)
				if err != nil {
					return nil, fmt.Errorf("failed to get checksum for input file: %v", err)
				}
			} else {
				thumb := make([]byte, thumbLen)
				_, err := exifReader.ReadAt(thumb, int64(thumbPos))
				if err != nil {
					return nil, fmt.Errorf("failed to get checksum for input file: %v", err)
				}
				_, err = h.Write(thumb)
				if err != nil {
					return nil, fmt.Errorf("failed to get checksum for input file: %v", err)
				}
			}
			dateTimeIntf, err := t.DecodeFirst(exif.DateTime)
			if err == nil {
				dateTime := dateTimeIntf.(string)
				t, err := time.Parse("2006:01:02 15:04:05", dateTime)
				if err == nil {
					captureTime = t
					prefix = t.Format("2006/01/02/15-04-05_")
				}
			}
		}
	}
	fd.Seek(0, io.SeekStart)
	checksum := fmt.Sprintf("%x", h.Sum(nil))
	basename := strings.TrimSuffix(filepath.Base(path), ext)
	return &Picture{
		Extension:   filepath.Ext(path),
		Basename:    basename,
		CaptureTime: captureTime,
		Name:        prefix + basename + "_" + checksum,
		Checksum:    checksum,
		Reader:      fd,
		original:    path,
	}, nil
}
