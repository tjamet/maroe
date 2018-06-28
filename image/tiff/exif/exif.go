package exif

import "github.com/tjamet/maroe/image/tiff"

const (
	exifPointer    = uint16(0x8769)
	gpsPointer     = 0x8825
	interopPointer = 0xA005
)

// Source: https://github.com/rwcarlsen/goexif/blob/go1/exif/fields.go
// Primary EXIF fields
const (
	ImageWidth                 = "ImageWidth"
	ImageLength                = "ImageLength" // Image height called Length by EXIF spec
	BitsPerSample              = "BitsPerSample"
	Compression                = "Compression"
	PhotometricInterpretation  = "PhotometricInterpretation"
	Orientation                = "Orientation"
	SamplesPerPixel            = "SamplesPerPixel"
	RowsPerStrip               = "RowsPerStrip"
	PlanarConfiguration        = "PlanarConfiguration"
	YCbCrSubSampling           = "YCbCrSubSampling"
	YCbCrPositioning           = "YCbCrPositioning"
	XResolution                = "XResolution"
	YResolution                = "YResolution"
	ResolutionUnit             = "ResolutionUnit"
	DateTime                   = "DateTime"
	ImageDescription           = "ImageDescription"
	Make                       = "Make"
	Model                      = "Model"
	Software                   = "Software"
	Artist                     = "Artist"
	Copyright                  = "Copyright"
	ExifIFDPointer             = "ExifIFDPointer"
	GPSInfoIFDPointer          = "GPSInfoIFDPointer"
	InteroperabilityIFDPointer = "InteroperabilityIFDPointer"
	ExifVersion                = "ExifVersion"
	FlashpixVersion            = "FlashpixVersion"
	ColorSpace                 = "ColorSpace"
	ComponentsConfiguration    = "ComponentsConfiguration"
	CompressedBitsPerPixel     = "CompressedBitsPerPixel"
	PixelXDimension            = "PixelXDimension"
	PixelYDimension            = "PixelYDimension"
	MakerNote                  = "MakerNote"
	UserComment                = "UserComment"
	RelatedSoundFile           = "RelatedSoundFile"
	DateTimeOriginal           = "DateTimeOriginal"
	DateTimeDigitized          = "DateTimeDigitized"
	SubSecTime                 = "SubSecTime"
	SubSecTimeOriginal         = "SubSecTimeOriginal"
	SubSecTimeDigitized        = "SubSecTimeDigitized"
	ImageUniqueID              = "ImageUniqueID"
	ExposureTime               = "ExposureTime"
	FNumber                    = "FNumber"
	ExposureProgram            = "ExposureProgram"
	SpectralSensitivity        = "SpectralSensitivity"
	ISOSpeedRatings            = "ISOSpeedRatings"
	OECF                       = "OECF"
	ShutterSpeedValue          = "ShutterSpeedValue"
	ApertureValue              = "ApertureValue"
	BrightnessValue            = "BrightnessValue"
	ExposureBiasValue          = "ExposureBiasValue"
	MaxApertureValue           = "MaxApertureValue"
	SubjectDistance            = "SubjectDistance"
	MeteringMode               = "MeteringMode"
	LightSource                = "LightSource"
	Flash                      = "Flash"
	FocalLength                = "FocalLength"
	SubjectArea                = "SubjectArea"
	FlashEnergy                = "FlashEnergy"
	SpatialFrequencyResponse   = "SpatialFrequencyResponse"
	FocalPlaneXResolution      = "FocalPlaneXResolution"
	FocalPlaneYResolution      = "FocalPlaneYResolution"
	FocalPlaneResolutionUnit   = "FocalPlaneResolutionUnit"
	SubjectLocation            = "SubjectLocation"
	ExposureIndex              = "ExposureIndex"
	SensingMethod              = "SensingMethod"
	FileSource                 = "FileSource"
	SceneType                  = "SceneType"
	CFAPattern                 = "CFAPattern"
	CustomRendered             = "CustomRendered"
	ExposureMode               = "ExposureMode"
	WhiteBalance               = "WhiteBalance"
	DigitalZoomRatio           = "DigitalZoomRatio"
	FocalLengthIn35mmFilm      = "FocalLengthIn35mmFilm"
	SceneCaptureType           = "SceneCaptureType"
	GainControl                = "GainControl"
	Contrast                   = "Contrast"
	Saturation                 = "Saturation"
	Sharpness                  = "Sharpness"
	DeviceSettingDescription   = "DeviceSettingDescription"
	SubjectDistanceRange       = "SubjectDistanceRange"
	LensMake                   = "LensMake"
	LensModel                  = "LensModel"
	CameraSerialNumber         = "CameraSerialNumber"
)

// thumbnail fields
const (
	ThumbnailOffset    = "ThumbnailOffset" // offset to thumb jpeg SOI
	ThumbnailLength    = "ThumbnailLength" // byte length of thumb
	ThumbnailTIFF      = "ThumbnailTIFF"
	PreviewImage       = "PreviewImage"
	PreviewImageStart  = "PreviewImageStart"
	PreviewImageLength = "PreviewImageLength"
	ThumbnailFormat    = "ThumbnailFormat"
)

// GPS fields
const (
	GPSVersionID        = "GPSVersionID"
	GPSLatitudeRef      = "GPSLatitudeRef"
	GPSLatitude         = "GPSLatitude"
	GPSLongitudeRef     = "GPSLongitudeRef"
	GPSLongitude        = "GPSLongitude"
	GPSAltitudeRef      = "GPSAltitudeRef"
	GPSAltitude         = "GPSAltitude"
	GPSTimeStamp        = "GPSTimeStamp"
	GPSSatelites        = "GPSSatelites"
	GPSStatus           = "GPSStatus"
	GPSMeasureMode      = "GPSMeasureMode"
	GPSDOP              = "GPSDOP"
	GPSSpeedRef         = "GPSSpeedRef"
	GPSSpeed            = "GPSSpeed"
	GPSTrackRef         = "GPSTrackRef"
	GPSTrack            = "GPSTrack"
	GPSImgDirectionRef  = "GPSImgDirectionRef"
	GPSImgDirection     = "GPSImgDirection"
	GPSMapDatum         = "GPSMapDatum"
	GPSDestLatitudeRef  = "GPSDestLatitudeRef"
	GPSDestLatitude     = "GPSDestLatitude"
	GPSDestLongitudeRef = "GPSDestLongitudeRef"
	GPSDestLongitude    = "GPSDestLongitude"
	GPSDestBearingRef   = "GPSDestBearingRef"
	GPSDestBearing      = "GPSDestBearing"
	GPSDestDistanceRef  = "GPSDestDistanceRef"
	GPSDestDistance     = "GPSDestDistance"
	GPSProcessingMethod = "GPSProcessingMethod"
	GPSAreaInformation  = "GPSAreaInformation"
	GPSDateStamp        = "GPSDateStamp"
	GPSDifferential     = "GPSDifferential"
)

// interoperability fields
const (
	InteroperabilityIndex = "InteroperabilityIndex"
)

const (
	ImageNumber = "ImageNumber"
	ImageID     = "ImageID"
)

const (
	DefaultCropOrigin = "DefaultCropOrigin"
	DefaultCropSize   = "DefaultCropSize"

	CFARepeatPatternDim = "CFARepeatPatternDim"
	CFAPattern2         = "CFAPattern2"
	BlackLevel          = "BlackLevel"
	WhiteLevel          = "WhiteLevel"
	ActiveArea          = "ActiveArea"
	TIFFEPStandardID    = "TIFF-EPStandardID"
	// XMP Tags
	ApplicationNotes = "ApplicationNotes"
)

func init() {
	tiff.RegisterSubIFD(exifPointer, "ExifIFD")
	/////////////////////////////////////
	////////// IFD 0 ////////////////////
	/////////////////////////////////////

	// image data structure for the thumbnail
	tiff.RegisterTag(0x0100, ImageWidth)
	tiff.RegisterTag(0x0101, ImageLength)
	tiff.RegisterTag(0x0102, BitsPerSample)
	tiff.RegisterTag(0x0103, Compression)
	tiff.RegisterTag(0x0106, PhotometricInterpretation)
	tiff.RegisterTag(0x0112, Orientation)
	tiff.RegisterTag(0x0115, SamplesPerPixel)
	tiff.RegisterTag(0x0116, RowsPerStrip)
	tiff.RegisterTag(0x011C, PlanarConfiguration)
	tiff.RegisterTag(0x0212, YCbCrSubSampling)
	tiff.RegisterTag(0x0213, YCbCrPositioning)
	tiff.RegisterTag(0x011A, XResolution)
	tiff.RegisterTag(0x011B, YResolution)
	tiff.RegisterTag(0x0128, ResolutionUnit)

	// Other tags
	tiff.RegisterTag(0x0132, DateTime)
	tiff.RegisterTag(0x010E, ImageDescription)
	tiff.RegisterTag(0x010F, Make)
	tiff.RegisterTag(0x0110, Model)
	tiff.RegisterTag(0x0131, Software)
	tiff.RegisterTag(0x013B, Artist)
	tiff.RegisterTag(0x8298, Copyright)

	// private tags
	tiff.RegisterTag(exifPointer, ExifIFDPointer)

	/////////////////////////////////////
	////////// Exif sub IFD /////////////
	/////////////////////////////////////

	tiff.RegisterTag(gpsPointer, GPSInfoIFDPointer)
	tiff.RegisterTag(interopPointer, InteroperabilityIFDPointer)

	tiff.RegisterTag(0x9000, ExifVersion)
	tiff.RegisterTag(0xA000, FlashpixVersion)

	tiff.RegisterTag(0xA001, ColorSpace)

	tiff.RegisterTag(0x9101, ComponentsConfiguration)
	tiff.RegisterTag(0x9102, CompressedBitsPerPixel)
	tiff.RegisterTag(0xA002, PixelXDimension)
	tiff.RegisterTag(0xA003, PixelYDimension)

	tiff.RegisterTag(0x927C, MakerNote)
	tiff.RegisterTag(0x9286, UserComment)

	tiff.RegisterTag(0xA004, RelatedSoundFile)
	tiff.RegisterTag(0x9003, DateTimeOriginal)
	tiff.RegisterTag(0x9004, DateTimeDigitized)
	tiff.RegisterTag(0x9290, SubSecTime)
	tiff.RegisterTag(0x9291, SubSecTimeOriginal)
	tiff.RegisterTag(0x9292, SubSecTimeDigitized)

	tiff.RegisterTag(0xA420, ImageUniqueID)

	// picture conditions
	tiff.RegisterTag(0x829A, ExposureTime)
	tiff.RegisterTag(0x829D, FNumber)
	tiff.RegisterTag(0x8822, ExposureProgram)
	tiff.RegisterTag(0x8824, SpectralSensitivity)
	tiff.RegisterTag(0x8827, ISOSpeedRatings)
	tiff.RegisterTag(0x8828, OECF)
	tiff.RegisterTag(0x9201, ShutterSpeedValue)
	tiff.RegisterTag(0x9202, ApertureValue)
	tiff.RegisterTag(0x9203, BrightnessValue)
	tiff.RegisterTag(0x9204, ExposureBiasValue)
	tiff.RegisterTag(0x9205, MaxApertureValue)
	tiff.RegisterTag(0x9206, SubjectDistance)
	tiff.RegisterTag(0x9207, MeteringMode)
	tiff.RegisterTag(0x9208, LightSource)
	tiff.RegisterTag(0x9209, Flash)
	tiff.RegisterTag(0x920A, FocalLength)
	tiff.RegisterTag(0x9214, SubjectArea)
	tiff.RegisterTag(0xA20B, FlashEnergy)
	tiff.RegisterTag(0xA20C, SpatialFrequencyResponse)
	tiff.RegisterTag(0xA20E, FocalPlaneXResolution)
	tiff.RegisterTag(0xA20F, FocalPlaneYResolution)
	tiff.RegisterTag(0xA210, FocalPlaneResolutionUnit)
	tiff.RegisterTag(0xA214, SubjectLocation)
	tiff.RegisterTag(0xA215, ExposureIndex)
	tiff.RegisterTag(0xA217, SensingMethod)
	tiff.RegisterTag(0xA300, FileSource)
	tiff.RegisterTag(0xA301, SceneType)
	tiff.RegisterTag(0xA302, CFAPattern)
	tiff.RegisterTag(0xA401, CustomRendered)
	tiff.RegisterTag(0xA402, ExposureMode)
	tiff.RegisterTag(0xA403, WhiteBalance)
	tiff.RegisterTag(0xA404, DigitalZoomRatio)
	tiff.RegisterTag(0xA405, FocalLengthIn35mmFilm)
	tiff.RegisterTag(0xA406, SceneCaptureType)
	tiff.RegisterTag(0xA407, GainControl)
	tiff.RegisterTag(0xA408, Contrast)
	tiff.RegisterTag(0xA409, Saturation)
	tiff.RegisterTag(0xA40A, Sharpness)
	tiff.RegisterTag(0xA40B, DeviceSettingDescription)
	tiff.RegisterTag(0xA40C, SubjectDistanceRange)
	tiff.RegisterTag(0xA433, LensMake)
	tiff.RegisterTag(0xA434, LensModel)

	/////////////////////////////////////
	//// GPS sub-IFD ////////////////////
	/////////////////////////////////////
	tiff.RegisterTag(0x0, GPSVersionID)
	tiff.RegisterTag(0x1, GPSLatitudeRef)
	tiff.RegisterTag(0x2, GPSLatitude)
	tiff.RegisterTag(0x3, GPSLongitudeRef)
	tiff.RegisterTag(0x4, GPSLongitude)
	tiff.RegisterTag(0x5, GPSAltitudeRef)
	tiff.RegisterTag(0x6, GPSAltitude)
	tiff.RegisterTag(0x7, GPSTimeStamp)
	tiff.RegisterTag(0x8, GPSSatelites)
	tiff.RegisterTag(0x9, GPSStatus)
	tiff.RegisterTag(0xA, GPSMeasureMode)
	tiff.RegisterTag(0xB, GPSDOP)
	tiff.RegisterTag(0xC, GPSSpeedRef)
	tiff.RegisterTag(0xD, GPSSpeed)
	tiff.RegisterTag(0xE, GPSTrackRef)
	tiff.RegisterTag(0xF, GPSTrack)
	tiff.RegisterTag(0x10, GPSImgDirectionRef)
	tiff.RegisterTag(0x11, GPSImgDirection)
	tiff.RegisterTag(0x12, GPSMapDatum)
	tiff.RegisterTag(0x13, GPSDestLatitudeRef)
	tiff.RegisterTag(0x14, GPSDestLatitude)
	tiff.RegisterTag(0x15, GPSDestLongitudeRef)
	tiff.RegisterTag(0x16, GPSDestLongitude)
	tiff.RegisterTag(0x17, GPSDestBearingRef)
	tiff.RegisterTag(0x18, GPSDestBearing)
	tiff.RegisterTag(0x19, GPSDestDistanceRef)
	tiff.RegisterTag(0x1A, GPSDestDistance)
	tiff.RegisterTag(0x1B, GPSProcessingMethod)
	tiff.RegisterTag(0x1C, GPSAreaInformation)
	tiff.RegisterTag(0x1D, GPSDateStamp)
	tiff.RegisterTag(0x1E, GPSDifferential)

	/////////////////////////////////////
	//// Interoperability sub-IFD ///////
	/////////////////////////////////////
	tiff.RegisterTag(0x1, InteroperabilityIndex)

	tiff.RegisterTag(0x0201, ThumbnailOffset)
	tiff.RegisterTag(0x0202, ThumbnailLength)
	tiff.RegisterTag(0x5012, ThumbnailFormat)

	tiff.RegisterTag(0x203, ThumbnailTIFF)
	tiff.RegisterTag(0x204, PreviewImage)

	tiff.RegisterTag(0x111, PreviewImageStart)
	tiff.RegisterTag(0x117, PreviewImageLength)

	tiff.RegisterTag(0x9211, ImageNumber)
	tiff.RegisterTag(0x800d, ImageID)

	tiff.RegisterTag(0xc62f, CameraSerialNumber)

	tiff.RegisterTag(0xc61f, DefaultCropOrigin)
	tiff.RegisterTag(0xc620, DefaultCropSize)

	tiff.RegisterTag(0x828d, CFARepeatPatternDim)
	tiff.RegisterTag(0x828e, CFAPattern2)

	tiff.RegisterTag(0xc61a, BlackLevel)
	tiff.RegisterTag(0xc61d, WhiteLevel)
	tiff.RegisterTag(0xc68d, ActiveArea)

	tiff.RegisterTag(0x9216, TIFFEPStandardID)

	tiff.RegisterTag(0x02bc, ApplicationNotes)

}
