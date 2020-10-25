//
// Created by Hugo Nijhuis-Mekkelholt on 25/10/2020.
//

#include "Metal/Metal.h"
#include "texture.h"

void mtl_Texture_Release(void *texture) {
    [(id <MTLTexture>) texture release];
}
