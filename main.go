package main

import (
	"log"

	"github.com/levantocode/tibia-go-otbm-parser/otbm"
)

func main() {
	_, err := otbm.NewOTBMReader("resources/tfs.otbm")
	if err != nil {
		log.Fatalf("Failed to load OTBM file: %v", err)
	}
}
