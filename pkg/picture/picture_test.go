package picture_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjamet/maroe/pkg/picture"
	tools "github.com/tjamet/maroe/pkg/test-tools"
)

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

func TestNewPictureFromFilePath(t *testing.T) {
	p, err := picture.NewPictureFromFilePath("../lister/test/resources/non-supported.EXT")
	assert.NoError(t, err)
	t.Run("Has lower case extension", func(t *testing.T) { assert.Equal(t, ".EXT", p.Extension) })
	t.Run("Has valid checksum", func(t *testing.T) {
		assert.Equal(t, "non-supported_6cebd3ce86795a7c5c080b460994730e7ce11686", p.Name)
	})
	t.Run("Returns a fully reabadle reader", func(t *testing.T) { assert.Equal(t, "some content\nsome other content", streamToString(p.Reader)) })
}

func testLoadImage(t testing.TB, p string) {
	t.Helper()
	_, err := picture.NewPictureFromFilePath(p)
	assert.NoError(t, err)
}

func BenchmarkNewPictureFromLosslessNEF(b *testing.B) {
	// http://www.rawsamples.ch/index.php/en/nikon
	path := tools.DownloadTestRAW(b, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_12bit_FX_LOSSLESS.NEF")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testLoadImage(b, path)
	}
}

func BenchmarkNewPictureFromLossyNEF(b *testing.B) {
	// http://www.rawsamples.ch/index.php/en/nikon
	path := tools.DownloadTestRAW(b, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_12bit_FX_LOSSY.NEF")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testLoadImage(b, path)
	}
}

func BenchmarkNewPictureFromUncompressedNEF(b *testing.B) {
	// http://www.rawsamples.ch/index.php/en/nikon
	path := tools.DownloadTestRAW(b, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_12bit_FX_UNCOMPRESSED.NEF")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testLoadImage(b, path)
	}
}

func BenchmarkNewPictureFromNikonTIF(b *testing.B) {
	// http://www.rawsamples.ch/index.php/en/nikon
	path := tools.DownloadTestRAW(b, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_L.TIFF")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testLoadImage(b, path)
	}
}

func BenchmarkNewPictureFromTIF(b *testing.B) {
	// https://file-examples.com/index.php/sample-images-download/sample-tiff-download/#google_vignette
	path := tools.DownloadTestRAW(b, "https://file-examples.com/wp-content/storage/2017/10/file_example_TIFF_10MB.tiff")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testLoadImage(b, path)
	}
}

func BenchmarkNewPictureFromDxOOneDNG(b *testing.B) {
	// https://www.imaging-resource.com/PRODS/dxo-one/dxo-one-field-test-part-i.htm
	path := tools.DownloadTestRAW(b, "https://www.imaging-resource.com/PRODS/dxo-one/FULLRES/YDXO_0406.DNG")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testLoadImage(b, path)
	}
}

func BenchmarkNewPictureFromLeicaDNG(b *testing.B) {
	// https://www.kenrockwell.com/leica/m9/sample-photos-3.htm
	path := tools.DownloadTestRAW(b, "https://www.kenrockwell.com/trips/2009-10/images/L1004220.DNG")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testLoadImage(b, path)
	}
}
