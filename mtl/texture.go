// +build darwin

package mtl

// #include "mtl/texture.h"
import "C"
import "unsafe"

type TextureHandle Handle

type Texture struct {
	Handle TextureHandle
}

func (t Texture) Release() {
	C.mtl_Texture_Release(unsafe.Pointer(t.Handle))
}
