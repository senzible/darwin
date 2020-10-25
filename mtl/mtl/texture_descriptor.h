//
// Created by Hugo Nijhuis-Mekkelholt on 24/10/2020.
//

#ifndef MTL_TEXTURE_DESCRIPTOR_H
#define MTL_TEXTURE_DESCRIPTOR_H

void mtl_TextureDescriptor_Release(void *textureDescriptor);

void *mtl_Texture2DDescriptorWithPixelFormat(unsigned short pixelFormat, unsigned long width, unsigned long height,
                                             signed char mipmapped);

void mtl_TextureDescriptor_SetUsage(void *textureDescriptor, unsigned short usage);

#endif //MTL_TEXTURE_DESCRIPTOR_H
