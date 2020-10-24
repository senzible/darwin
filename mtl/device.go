// +build darwin

package mtl

// #include "mtl/device.h"
import "C"
import "unsafe"

type DeviceHandle Handle

type Device struct {
	Handle DeviceHandle
}

func (d Device) Name() string {
	return C.GoString(C.mtl_Device_Name(unsafe.Pointer(d.Handle)))
}

func (d Device) SupportsFamily(family MTLGPUFamily) bool {
	return C.mtl_Device_SupportsFamily(unsafe.Pointer(d.Handle), C.ushort(family)) != 0
}
