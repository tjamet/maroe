package writer_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjamet/maroe/pkg/picture"
	"github.com/tjamet/maroe/pkg/writer"
)

func TestWrite(t *testing.T) {
	os.Remove("test/written/pfx/abcdef.ext")
	p := &picture.Picture{
		Extension: ".ext",
		Name:      "pfx/abcdef",
		Reader:    bytes.NewBufferString("some\ncontent"),
	}
	path, err := writer.Write(p, "test/written")
	t.Run("Write does not report error", func(t *testing.T) { assert.NoError(t, err) })
	t.Run("new picture path is returned", func(t *testing.T) { assert.Equal(t, "test/written/pfx/abcdef.ext", path) })
	t.Run("picture is written", func(t *testing.T) { assert.FileExists(t, "test/written/pfx/abcdef.ext") })

	t.Run("Writing the same file again", func(t *testing.T) {
		path, err := writer.Write(p, "test/written")
		t.Run("Write reports error", func(t *testing.T) { assert.Error(t, err) })
		t.Run("new picture path is returned", func(t *testing.T) { assert.Equal(t, "test/written/pfx/abcdef.ext", path) })
		t.Run("picture is not overwritten", func(t *testing.T) { assert.FileExists(t, "test/written/pfx/abcdef.ext") })
	})
}
