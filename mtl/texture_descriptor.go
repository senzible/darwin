// +build darwin

package mtl

// #include <stdlib.h>
// #include "mtl/texture_descriptor.h"
import "C"
import (
	"errors"
	"unsafe"
)

type TextureDescriptorHandle Handle

type TextureDescriptor struct {
	Handle TextureDescriptorHandle
}

// Texture2DDescriptorWithPixelFormat Creates a texture descriptor object for a 2D texture.
// Make sure to call TextureDescriptor.Release() after you're done with it.
func Texture2DDescriptorWithPixelFormat(pixelFormat PixelFormat, width int, height int, mipmapped bool) (TextureDescriptor, error) {
	td := C.mtl_Texture2DDescriptorWithPixelFormat(C.ushort(pixelFormat), C.ulong(width), C.ulong(height), C.schar(boolToInt(mipmapped)))

	if td == nil {
		return TextureDescriptor{}, errors.New("Texture2DDescriptorWithPixelFormat failed")
	}

	return TextureDescriptor{
		Handle: TextureDescriptorHandle(td),
	}, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Release the TextureDescriptor memory.
func (td TextureDescriptor) Release() {
	C.mtl_TextureDescriptor_Release(unsafe.Pointer(td.Handle))
}

func (td TextureDescriptor) SetUsage(usage TextureUsage) {
	C.mtl_TextureDescriptor_SetUsage(unsafe.Pointer(td.Handle), C.ushort(usage))
}
