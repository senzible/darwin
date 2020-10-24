//
// Created by Hugo Nijhuis-Mekkelholt on 24/10/2020.
//

#ifndef MTL_DEVICE_H
#define MTL_DEVICE_H

const char *mtl_Device_Name(void *device);

signed char mtl_Device_SupportsFamily(void *device, unsigned short gpuFamily);

#endif //MTL_DEVICE_H
