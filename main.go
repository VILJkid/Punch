package main

import (
	"log"

	"github.com/VILJkid/go-family-tree/family"
)

func main() {
	// entrypoint
	if err := family.RunCommand(); err != nil {
		log.Println(err)
		return
	}
	log.Println("Execution successful")
}
