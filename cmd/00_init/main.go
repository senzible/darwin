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

	log.Printf("Metal Device Created: %s\n", d.Name())

	log.Printf("MTLGPUFamilyApple1: %t", d.SupportsFamily(mtl.MTLGPUFamilyApple1))
	log.Printf("MTLGPUFamilyApple2: %t", d.SupportsFamily(mtl.MTLGPUFamilyApple2))
	log.Printf("MTLGPUFamilyApple3: %t", d.SupportsFamily(mtl.MTLGPUFamilyApple3))
	log.Printf("MTLGPUFamilyApple4: %t", d.SupportsFamily(mtl.MTLGPUFamilyApple4))
	log.Printf("MTLGPUFamilyApple5: %t", d.SupportsFamily(mtl.MTLGPUFamilyApple5))
	log.Printf("MTLGPUFamilyApple6: %t", d.SupportsFamily(mtl.MTLGPUFamilyApple6))
	log.Printf("MTLGPUFamilyApple7: %t", d.SupportsFamily(mtl.MTLGPUFamilyApple7))

	log.Printf("MTLGPUFamilyMac1: %t", d.SupportsFamily(mtl.MTLGPUFamilyMac1))
	log.Printf("MTLGPUFamilyMac2: %t", d.SupportsFamily(mtl.MTLGPUFamilyMac2))

	log.Printf("MTLGPUFamilyCommon1: %t", d.SupportsFamily(mtl.MTLGPUFamilyCommon1))
	log.Printf("MTLGPUFamilyCommon2: %t", d.SupportsFamily(mtl.MTLGPUFamilyCommon2))
	log.Printf("MTLGPUFamilyCommon3: %t", d.SupportsFamily(mtl.MTLGPUFamilyCommon3))

	log.Printf("MTLGPUFamilyMacCatalyst1: %t", d.SupportsFamily(mtl.MTLGPUFamilyMacCatalyst1))
	log.Printf("MTLGPUFamilyMacCatalyst2: %t", d.SupportsFamily(mtl.MTLGPUFamilyMacCatalyst2))

}
