package writer_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tjamet/maroe/pkg/picture"
	"github.com/tjamet/maroe/pkg/writer"
)

func TestWrite(t *testing.T) {
	fs := afero.NewMemMapFs()

	writer := writer.NewFSWriter(fs)
	os.Remove("test/written/pfx/abcdef.ext")
	p := &picture.Picture{
		Extension: ".ext",
		Name:      "pfx/abcdef",
		Reader:    bytes.NewBufferString("some\ncontent"),
	}
	path, err := writer.Write(p, "test/written")
	t.Run("Write does not report error", func(t *testing.T) { assert.NoError(t, err) })
	t.Run("new picture path is returned", func(t *testing.T) { assert.Equal(t, "test/written/pfx/abcdef.ext", path) })
	t.Run("picture is written", func(t *testing.T) {
		_, err := fs.Stat("test/written/pfx/abcdef.ext")
		assert.NoError(t, err)
	})

	t.Run("Writing the same file again", func(t *testing.T) {
		fd, err := fs.Create("test/written/pfx/abcdef.ext")
		require.NoError(t, err)
		_, err = fd.Write([]byte("old\ncontent"))
		require.NoError(t, err)
		require.NoError(t, fd.Close())
		path, err := writer.Write(p, "test/written")
		t.Run("Write reports error", func(t *testing.T) { assert.Error(t, err) })
		t.Run("new picture path is returned", func(t *testing.T) { assert.Equal(t, "test/written/pfx/abcdef.ext", path) })
		t.Run("picture is not overwritten", func(t *testing.T) {
			fd, err := fs.Open("test/written/pfx/abcdef.ext")
			assert.NoError(t, err)
			data, err := io.ReadAll(fd)
			require.NoError(t, err)
			assert.Equal(t, "old\ncontent", string(data))
		})
	})
}
