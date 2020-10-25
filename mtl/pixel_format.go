package mtl

// PixelFormat is the data formats that describe the organization and characteristics of individual pixels in a texture.
type PixelFormat uint16

//goland:noinspection SpellCheckingInspection
const (
	// PixelFormatRGBA8UNorm Ordinary format with four 8-bit normalized unsigned integer components in RGBA order.
	PixelFormatRGBA8UNorm PixelFormat = 70

	// PixelFormatBGRA8Unorm ordinary format with four 8-bit normalized unsigned integer components in BGRA order.
	PixelFormatBGRA8Unorm PixelFormat = 80
)
