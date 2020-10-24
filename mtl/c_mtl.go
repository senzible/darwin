// +build darwin

package mtl

/*
#cgo CFLAGS: -x objective-c -mmacosx-version-min=10.11
#cgo LDFLAGS: -framework Metal -framework Foundation -framework CoreGraphics
#include "mtl/mtl.m"
*/
import "C"