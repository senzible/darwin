//
// Created by Hugo Nijhuis-Mekkelholt on 24/10/2020.
//

#include "Metal/Metal.h"
#include "device.h"

const char *mtl_Device_Name(void *device) {
    return ((id <MTLDevice>) device).name.UTF8String;
}

signed char mtl_Device_SupportsFamily(void *device, unsigned short gpuFamily) {
    return [(id <MTLDevice>) device supportsFamily:gpuFamily];
}