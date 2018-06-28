package tiff_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tjamet/goraw"
	_ "github.com/tjamet/goraw/all"
	"github.com/tjamet/maroe/image/tiff"
	_ "github.com/tjamet/maroe/image/tiff/exif"
	tools "github.com/tjamet/maroe/pkg/test-tools"
)

func TestCANONrawscanonRAW_CANON_EOS_1DXCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_1DX.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2012:09:06 07:48:25")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon1dsm3RAW_CANON_1DSM3CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/1dsm3/RAW_CANON_1DSM3.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:10:18 13:44:32")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon1dsm2RAW_CANON_1DSM2CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/1dsm2/RAW_CANON_1DSM2.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2004:11:13 23:02:21")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_1DM4CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_1DM4.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2009:12:29 16:45:53")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon1dm3RAW_CANON_1DMARK3CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/1dm3/RAW_CANON_1DMARK3.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:07:14 18:02:56")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon1dm2nRAW_CANON_1DM2NCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/1dm2n/RAW_CANON_1DM2N.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:03:18 11:56:05")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon1dm2RAW_CANON_1DM2CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/1dm2/RAW_CANON_1DM2.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2005:10:29 16:14:44")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_5DSCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_5DS.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:07:23 16:52:30")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_5DMARK3CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_5DMARK3.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2012:03:08 16:08:39")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon5dm2RAW_CANON_5DMARK2_PREPRODCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/5dm2/RAW_CANON_5DMARK2_PREPROD.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:10:29 20:05:00")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon5dRAW_CANON_5D_ARGBCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/5d/RAW_CANON_5D_ARGB.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2006:01:15 19:04:48")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_6DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_6D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:12:31 16:08:56")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_7DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_7D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2009:10:09 14:18:45")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS70DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS70D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:13 12:05:21")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_60D_V108_HORIZONTALCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_60D_V108_HORIZONTAL.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2011:03:20 03:38:10")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_60D_V108_VERTICALCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_60D_V108_VERTICAL.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2011:03:20 03:39:09")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon50dRAW_CANON_50DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/50d/RAW_CANON_50D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:09:17 15:13:40")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon40dRAW_CANON_40D_RAW_V336643CCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_RAW_V336643C.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:11:16 13:00:47")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon40dRAW_CANON_40D_RAW_V105CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_RAW_V105.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:12:23 20:32:28")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon40dRAW_CANON_40D_RAW_V104CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_RAW_V104.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:11:06 11:24:20")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon40dRAW_CANON_40D_SRAW_V103CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_SRAW_V103.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:09:14 11:32:02")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon40dRAW_CANON_40D_RAW_V103CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/40d/RAW_CANON_40D_RAW_V103.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:09:14 11:32:59")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon30dRAW_CANON_30DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/30d/RAW_CANON_30D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:03:31 09:54:09")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon20dRAW_CANON_20DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/20d/RAW_CANON_20D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2005:06:18 11:43:54")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS1200DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS1200D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:10:17 09:13:48")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_1100DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_1100D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2012:12:15 19:02:01")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_1000DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_1000D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2010:07:23 11:00:15")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_760DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_760D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:10:27 09:42:25")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_700DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_700D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2013:12:18 11:32:12")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_600DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_600D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2011:07:14 19:38:47")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_550DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_550D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2010:03:03 21:50:57")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon450dRAW_CANON_450DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/450d/RAW_CANON_450D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:04:04 15:58:18")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon400dRAW_CANON_400D_ARGBCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/400d/RAW_CANON_400D_ARGB.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:02:21 18:24:36")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanon350dRAW_CANON_350DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/350d/RAW_CANON_350D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2005:04:04 16:58:39")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOS_100DCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS_100D.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:04:20 19:21:15")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOSM3CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS-M3.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:08:29 11:49:20")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_EOSMCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_EOS-M.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:07:26 05:53:10")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_G1XCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_G1X.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2013:11:23 16:23:30")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_G5XCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_G5X.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:11:15 23:15:18")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_G7XCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_G7X.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:04:17 14:47:35")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_G15CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_G15.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:03:18 12:10:12")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_G12CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_G12.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2011:02:21 14:09:39")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_G11CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_G11.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2010:06:27 15:34:54")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanong10RAW_CANON_G10CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/g10/RAW_CANON_G10.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:12:30 14:56:02")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanong9RAW_CANON_G9CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/g9/RAW_CANON_G9.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:10:17 19:55:27")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_SD450_CHDK_V097725DNG(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_SD450_CHDK_V0.9.7-725.DNG")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)

	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_S120CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_S120.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:07:06 16:47:00")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_SX100IS_CHDK_V089788DNG(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_SX100IS_CHDK_V0.8.9-788.DNG")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)

	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_S90CR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_S90.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2011:08:24 14:41:04")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_SX1_ISCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_SX1_IS.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2009:07:30 14:06:09")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_SX60HSCR2(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_SX60-HS.CR2")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:07:09 11:53:44")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_SX130IS_CHDKDNG(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_SX130IS_CHDK.DNG")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)

	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_SX510HSDNG(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_SX510HS.DNG")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)

	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestCANONrawscanonRAW_CANON_POWERSHOT_A3200IS_CHDKDNG(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/canon/RAW_CANON_POWERSHOT_A3200IS_CHDK.DNG")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)

	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond3xRAW_NIKON_D3XNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d3x/RAW_NIKON_D3X.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:12:01 14:52:13")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond3RAW_NIKON_D3NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d3/RAW_NIKON_D3.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:01:24 14:00:59")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond2xRAW_NIKON_D2X_SRGBNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d2x/RAW_NIKON_D2X_SRGB.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2005:05:13 23:26:25")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond1xRAW_NIKON_D1XNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d1x/RAW_NIKON_D1X.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2006:12:31 10:32:29")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond1RAW_NIKON_D1NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d1/RAW_NIKON_D1.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2000:11:19 13:01:50")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_DS200NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_DS200.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2013:06:19 20:31:15")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D7100NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D7100.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2013:09:03 09:45:16")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D5300NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D5300.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:04:02 11:39:52")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D5200NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D5200.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2013:11:27 09:21:07")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D5100NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D5100.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2012:05:04 16:27:29")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D3200NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D3200.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2013:09:04 13:11:48")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D3100NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D3100.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:07:28 20:40:46")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_12bit_FX_LOSSLESSNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_12bit_FX_LOSSLESS.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:31:54")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_12bit_FX_LOSSYNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_12bit_FX_LOSSY.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:32:27")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_12bit_FX_UNCOMPRESSEDNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_12bit_FX_UNCOMPRESSED.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:31:13")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_14bit_12x30x20_UNCOMPRESSEDNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_14bit_1.2x-30x20_UNCOMPRESSED.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:35:34")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_14bit_5x430x24_UNCOMPRESSEDNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_14bit_5x4-30x24_UNCOMPRESSED.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:38:04")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_14bit_DX24x16_UNCOMPRESSEDNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_14bit_DX-24x16_UNCOMPRESSED.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:37:47")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_14bit_FX_LOSSLESSNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_14bit_FX_LOSSLESS.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:24:26")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_14bit_FX_LOSSYNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_14bit_FX_LOSSY.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:27:35")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D800_14bit_FX_UNCOMPRESSEDNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D800_14bit_FX_UNCOMPRESSED.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:15 10:27:14")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D750NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D750.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:06:02 09:56:01")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond700RAW_NIKON_D700NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d700/RAW_NIKON_D700.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:10:25 14:41:48")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond300RAW_NIKON_D300NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d300/RAW_NIKON_D300.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:09:22 12:02:39")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D600NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D600.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:08:09 05:22:17")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_D300SNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_D300S.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2009:12:16 21:07:45")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond200RAW_NIKON_D200_SRGBNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d200/RAW_NIKON_D200_SRGB.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:04:01 19:31:17")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond100RAW_NIKON_D100NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d100/RAW_NIKON_D100.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:09:22 10:36:36")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond90RAW_NIKON_D90NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d90/RAW_NIKON_D90.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2009:02:10 19:47:07")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond80RAW_NIKON_D80_SRGBNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d80/RAW_NIKON_D80_SRGB.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:02:06 15:09:00")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond70RAW_NIKON_D70NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d70/RAW_NIKON_D70.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2004:06:26 14:16:26")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond50RAW_NIKON_D50NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d50/RAW_NIKON_D50.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:07:15 11:48:21")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikond40RAW_NIKON_D40_SRGBNEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/d40/RAW_NIKON_D40_SRGB.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:04:01 17:02:53")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikone5400RAW_NIKON_E5400NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/e5400/RAW_NIKON_E5400.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:07:04 19:02:37")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_COOLPIX_P7100NRW(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_COOLPIX_P7100.NRW")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2011:01:01 00:51:56")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_P7000NRW(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_P7000.NRW")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:08:08 21:56:44")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonp6000RAW_NIKON_P6000_GPSNRW(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/p6000/RAW_NIKON_P6000_GPS.NRW")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:12:14 13:55:22")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_COOLPIX_P340NRW(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_COOLPIX_P340.NRW")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:05:26 14:36:54")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON_1S2NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON_1S2.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:01:03 09:27:59")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestNIKONrawsnikonRAW_NIKON1_V1NEF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/nikon/RAW_NIKON1_V1.NEF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:06:26 16:12:36")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujie550RAW_FUJI_E550RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/e550/RAW_FUJI_E550.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:06:01 14:29:52")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujie900RAW_FUJI_E900RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/e900/RAW_FUJI_E900.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:05:20 11:26:08")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_F600XRRAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_F600XR.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:06:01 10:43:59")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujif700RAW_FUJI_F700RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/f700/RAW_FUJI_F700.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:01:06 14:20:22")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_FINPIX_F900EXRRAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_FINPIX_F900EXR.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2013:11:22 11:16:17")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_S9600RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_S9600.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2009:06:14 09:35:18")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujis9500RAW_FUJI_S9500_SRGBRAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/s9500/RAW_FUJI_S9500_SRGB.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2006:12:31 11:42:35")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujis6500fdRAW_FUJI_S6500FDRAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/s6500fd/RAW_FUJI_S6500FD.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:06:11 19:56:35")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujis5600RAW_FUJI_S5600S5200RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/s5600/RAW_FUJI_S5600S5200.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:04:22 10:47:03")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujis5000RAW_FUJI_S5000RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/s5000/RAW_FUJI_S5000.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2008:06:22 03:48:05")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_FINEPIX_S200EXRRAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_FINEPIX_S200EXR.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2010:07:07 05:12:03")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujis5proRAW_FUJI_S5PRO_V106RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/s5pro/RAW_FUJI_S5PRO_V106.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:05:27 13:55:17")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujis3proRAW_FUJI_S3PRORAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/s3pro/RAW_FUJI_S3PRO.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2007:06:03 18:28:41")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujis2proRAW_FUJI_S2PRORAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/s2pro/RAW_FUJI_S2PRO.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2004:03:06 16:28:37")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJIFILM_X100SRAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJIFILM_X100S.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:05:30 14:37:20")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_FINEPIX_X100RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_FINEPIX_X100.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2011:05:23 11:51:28")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_X20RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_X20.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:06:30 17:41:41")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_X10RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_X10.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:11:08 12:58:29")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_XQ1RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_XQ1.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:07:25 08:59:13")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_XPRO1RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_XPRO1.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:09:08 10:15:45")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_XA2RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_X-A2.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:07:19 12:54:32")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_XE1RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_X-E1.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:01:11 15:40:05")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUIJI_XE2RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUIJI_X-E2.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2014:10:12 17:28:29")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_XT10RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_X-T10.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:06:24 16:32:05")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJI_FINEPIX_HS10RAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJI_FINEPIX_HS10.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2010:07:16 09:43:55")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}

func TestFUJIrawsfujiRAW_FUJIFILM_HS20EXRRAF(t *testing.T) {
	path := tools.DownloadTestRAW(t, "http://www.rawsamples.ch/raws/fuji/RAW_FUJIFILM_HS20EXR.RAF")
	image, err := os.Open(path)
	assert.NoError(t, err)
	raw, err := goraw.New(image)
	require.NoError(t, err)
	require.NotNil(t, raw)
	exifReader, err := raw.ExifReaderAt()
	require.NoError(t, err)
	tiffImage, err := tiff.Read(exifReader)
	assert.NoError(t, err)
	assert.NotNil(t, tiffImage)
	testCanReadValue(t, tiffImage, 0x0132, "2015:09:26 16:37:50")
	if t.Failed() {
		tiffImage.DescribeTagAndValues()
	}
}
