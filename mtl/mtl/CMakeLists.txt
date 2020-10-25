cmake_minimum_required(VERSION 3.17)
project(mtl C)

set(CMAKE_C_STANDARD 99)

add_library(mtl mtl.m mtl.h device.m device.h texture_descriptor.m texture_descriptor.h texture.m texture.h)