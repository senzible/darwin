// +build darwin

package mtl

// #include <stdlib.h>
// #include "mtl/mtl.h"
import "C"
import (
	"unsafe"
)

type Handle unsafe.Pointer
