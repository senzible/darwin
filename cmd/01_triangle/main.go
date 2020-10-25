package main

import (
	"github.com/senzible/darwin/mtl"
	"log"
)

func main() {
	d, err := mtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}

	td, err := mtl.Texture2DDescriptorWithPixelFormat(mtl.PixelFormatRGBA8UNorm, 4, 4, false)
	if err != nil {
		log.Fatal(err)
	}
	defer td.Release()

	td.SetUsage(mtl.TextureUsageRenderTarget)

	t, err := d.NewTextureWithDescriptor(td)
	if err != nil {
		log.Fatal(err)
	}
	defer t.Release()
}
