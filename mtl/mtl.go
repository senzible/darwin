// +build darwin

package mtl

// #include "mtl/mtl.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Handle unsafe.Pointer

// CreateSystemDefaultDevice Returns a reference to the preferred default Metal device object.
func CreateSystemDefaultDevice() (Device, error) {
	d := C.mtl_CreateSystemDefaultDevice()
	if d == nil {
		return Device{}, errors.New("CreateSystemDefaultDevice failed")
	}

	return Device{
		Handle: DeviceHandle(d),
	}, nil
}
