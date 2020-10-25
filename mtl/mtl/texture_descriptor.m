//
// Created by Hugo Nijhuis-Mekkelholt on 24/10/2020.
//

#include "Metal/Metal.h"
#include "texture_descriptor.h"

void mtl_TextureDescriptor_Release(void *textureDescriptor) {
    [(MTLTextureDescriptor *) textureDescriptor release];
}

void *mtl_Texture2DDescriptorWithPixelFormat(unsigned short pixelFormat, unsigned long width, unsigned long height,
                                             signed char mipmapped) {
    MTLTextureDescriptor *textureDescriptor = [MTLTextureDescriptor texture2DDescriptorWithPixelFormat:(MTLPixelFormat) pixelFormat
                                                                                                 width:width
                                                                                                height:height
                                                                                             mipmapped:mipmapped];
    return textureDescriptor;
}

void mtl_TextureDescriptor_SetUsage(void *textureDescriptor, unsigned short usage) {
    ((MTLTextureDescriptor *) textureDescriptor).usage = (MTLTextureUsage) usage;
}