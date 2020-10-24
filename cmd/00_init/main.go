package main

import (
	"github.com/senzible/darwin/mtl"
	"log"
)

func main(){
	d, err :=mtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Metal Device Created: %s\n", d)
}