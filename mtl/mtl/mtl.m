//
// Created by Hugo Nijhuis-Mekkelholt on 24/10/2020.
//

#import <stddef.h>
#import <Metal/Metal.h>
#include "mtl.h"

void *mtl_CreateSystemDefaultDevice(){
    id <MTLDevice> device = MTLCreateSystemDefaultDevice();
    if(!device){
        return NULL;
    }
    return device;
}