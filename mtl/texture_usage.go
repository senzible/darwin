// +build darwin

package mtl

// TextureUsage an enumeration for the various options that determine how you can use a texture.
type TextureUsage uint8

const (
	TextureUsageUnknown         TextureUsage = 0x0000
	TextureUsageShaderRead      TextureUsage = 0x0001
	TextureUsageShaderWrite     TextureUsage = 0x0002
	TextureUsageRenderTarget    TextureUsage = 0x0004
	TextureUsagePixelFormatView TextureUsage = 0x0008
)
