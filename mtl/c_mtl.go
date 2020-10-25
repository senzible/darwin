// +build darwin

package mtl

/*
#cgo CFLAGS: -x objective-c -mmacosx-version-min=10.15
#cgo LDFLAGS: -framework Metal -framework Foundation -framework CoreGraphics
#include "mtl/mtl.m"
#include "mtl/device.m"
#include "mtl/texture_descriptor.m"
#include "mtl/texture.m"
*/
import "C"