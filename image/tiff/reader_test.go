package tiff_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjamet/maroe/image/tiff"
	tools "github.com/tjamet/maroe/pkg/test-tools"
)

func testCanReadValue[T any](t *testing.T, i *tiff.Tiff, code uint16, expected T) {
	t.Helper()
	t.Run(fmt.Sprintf("%d field can be read", code), func(t *testing.T) {
		rawValue, err := i.DecodeFirst(code)
		assert.NoError(t, err)
		assert.Equal(t, expected, rawValue)
	})
}

func BenchmarkLazyDecode(b *testing.B) {
	testFile := tools.DownloadTestRAW(b, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_5DS.CR2")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		fd, err := os.Open(testFile)
		if err != nil {
			b.Errorf("Failed to open test file %s", err.Error())
		}
		_, err = tiff.Read(fd)
		if err != nil {
			b.Errorf("Failed to decode test file %s", err.Error())
		}
		fd.Close()
	}
}
