// +build darwin

package mtl

// #include "mtl/mtl.h"
// #include "mtl/device.h"
import "C"
import (
	"errors"
	"unsafe"
)

type DeviceHandle Handle

type Device struct {
	Handle DeviceHandle
}

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

func (d Device) Name() string {
	return C.GoString(C.mtl_Device_Name(unsafe.Pointer(d.Handle)))
}

func (d Device) SupportsFamily(family MTLGPUFamily) bool {
	return C.mtl_Device_SupportsFamily(unsafe.Pointer(d.Handle), C.ushort(family)) != 0
}

// NewTextureWithDescriptor Creates a texture.
// Be sure to call Texture.Release() after you're done with it.
func (d Device) NewTextureWithDescriptor(descriptor TextureDescriptor) (Texture, error) {
	t := C.mtl_Device_NewTextureWithDescriptor(unsafe.Pointer(d.Handle), unsafe.Pointer(descriptor.Handle))
	if t == nil {
		return Texture{}, errors.New("NewTextureWithDescriptor failed")
	}

	return Texture{
		Handle: TextureHandle(t),
	}, nil
}
