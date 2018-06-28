package tiff

import (
	"encoding/binary"
	"io"
)

const (
	subIFDTag          = 0x14A
	SubIFD             = "SubIFD"
	SubfileType        = "SubfileType"
	OldSubfileType     = "OldSubfileType"
	ProcessingSoftware = "ProcessingSoftware"
)

// Most values taken from https://github.com/golang/image/blob/master/tiff/consts.go
const (
	// HeaderLittleEndian is the TIFF tag a file starts with to declare a little endian format
	HeaderLittleEndian = "II"
	// HeaderBigEndian is the TIFF tag a file starts with to declare a big endian format
	HeaderBigEndian  = "MM"
	tTagLen          = 12 // Length of an Tag entry in bytes.
	tTIFFMagicNumber = 42
	tNextIFDLen      = 4
)

// Data types (p. 14-16 of the spec).
const (
	dtByte     = 1
	dtASCII    = 2
	dtShort    = 3
	dtLong     = 4
	dtRational = 5
	dtIFD      = 13
)

// The length of one instance of each data type in bytes.
var lengths = map[uint16]uint32{
	dtByte:     1,
	dtASCII:    1,
	dtShort:    2,
	dtLong:     4,
	dtRational: 8,
	dtIFD:      4,
}

// Tags (see p. 28-41 of the spec).
const (
	tImageWidth                = 256
	tImageLength               = 257
	tBitsPerSample             = 258
	tCompression               = 259
	tPhotometricInterpretation = 262

	tStripOffsets    = 273
	tSamplesPerPixel = 277
	tRowsPerStrip    = 278
	tStripByteCounts = 279

	tTileWidth      = 322
	tTileLength     = 323
	tTileOffsets    = 324
	tTileByteCounts = 325

	tXResolution    = 282
	tYResolution    = 283
	tResolutionUnit = 296

	tPredictor    = 317
	tColorMap     = 320
	tExtraSamples = 338
	tSampleFormat = 339
)

// Compression types (defined in various places in the spec and supplements).
const (
	cNone       = 1
	cCCITT      = 2
	cG3         = 3 // Group 3 Fax.
	cG4         = 4 // Group 4 Fax.
	cLZW        = 5
	cJPEGOld    = 6 // Superseded by cJPEG.
	cJPEG       = 7
	cDeflate    = 8 // zlib compression.
	cPackBits   = 32773
	cDeflateOld = 32946 // Superseded by cDeflate.
)

// Photometric interpretation values (see p. 37 of the spec).
const (
	pWhiteIsZero = 0
	pBlackIsZero = 1
	pRGB         = 2
	pPaletted    = 3
	pTransMask   = 4 // transparency mask
	pCMYK        = 5
	pYCbCr       = 6
	pCIELab      = 8
)

// Values for the tPredictor tag (page 64-65 of the spec).
const (
	prNone       = 1
	prHorizontal = 2
)

// Values for the tResolutionUnit tag (page 18).
const (
	resNone    = 1
	resPerInch = 2 // Dots per inch.
	resPerCM   = 3 // Dots per centimeter.
)

// Tiff holds the accessors for a tiff based image (including canon CR2 files)
type Tiff struct {
	ByteOrder binary.ByteOrder
	ifds      []*tIFD
	reader    io.ReaderAt
}

type tIFD struct {
	offset uint32
	tag    []*tag
	name   string
}

type tag struct {
	code     uint16
	dataType uint16
	count    uint32
	ptr      []byte
	subIFDs  [][]*tIFD
}

func init() {
	RegisterTag(subIFDTag, SubIFD)

	RegisterTag(0xfe, SubfileType)
	RegisterTag(0xff, OldSubfileType)
	RegisterTag(0x000b, ProcessingSoftware)

	RegisterTag(0xc634, DNGAdobeData)
	RegisterTag(0xc612, DNGVersion)
	RegisterTag(0xc613, DNGBackwardVersion)
	RegisterTag(0xc614, UniqueCameraModel)
	RegisterTag(0xc615, LocalizedCameraModel)
	RegisterTag(0xc617, CFAPlaneColor)
	RegisterTag(0xc618, CFALayout)

	RegisterTag(0xc621, ColorMatrix1)
	RegisterTag(0xc622, ColorMatrix2)
	RegisterTag(0xc623, CameraCalibration1)
	RegisterTag(0xc624, CameraCalibration2)
	RegisterTag(0xc625, ReductionMatrix1)
	RegisterTag(0xc626, ReductionMatrix2)
	RegisterTag(0xc627, AnalogBalance)
	RegisterTag(0xc628, AsShotNeutral)
	RegisterTag(0xc629, AsShotWhiteXY)
	RegisterTag(0xc62a, BaselineExposure)
	RegisterTag(0xc62b, BaselineNoise)
	RegisterTag(0xc62c, BaselineSharpness)
	RegisterTag(0xc62d, BayerGreenSplit)
	RegisterTag(0xc62e, LinearResponseLimit)
	RegisterTag(0xc630, DNGLensInfo)
	RegisterTag(0xc631, ChromaBlurRadius)
	RegisterTag(0xc632, AntiAliasStrength)
	RegisterTag(0xc633, ShadowScale)
	RegisterTag(0xc635, MakerNoteSafety)
	RegisterTag(0xc65a, CalibrationIlluminant1)
	RegisterTag(0xc65b, CalibrationIlluminant2)
}
