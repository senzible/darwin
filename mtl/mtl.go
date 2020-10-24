// +build darwin

package mtl

// #include "mtl/mtl.h"
import "C"
import (
	"errors"
	"unsafe"
)

// CreateSystemDefaultDevice Returns a reference to the preferred default Metal device object.
func CreateSystemDefaultDevice() (unsafe.Pointer, error) {
	d := C.mtl_CreateSystemDefaultDevice()
	if d == nil {
		return nil, errors.New("CreateSystemDefaultDevice failed")
	}

	return d, nil
}
