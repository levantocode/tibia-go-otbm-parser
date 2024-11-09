package otbm

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

// Initializes a new OTBMReader.
func NewOTBMReader(filePath string) (*OTBMReader, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &OTBMReader{data: bytes.NewReader(file)}, nil
}

// Loads the map file and processes each node.
func (otbmReader *OTBMReader) StartReadingBytes() error {
	if ok, _ := otbmReader.hasValidMagicIdentifier(); !ok {
		otbmReader.DumpFileContents()
		fmt.Println("Invalid OTBM Magic Identifier detected.")
		return errors.New("invalid OTBM Magic Identifier")
	}

	if ok, _ := otbmReader.ReadHeader(); !ok {
		fmt.Println("Invalid OTBM Header detected.")
		otbmReader.DumpFileContents()
		return errors.New("invalid OTBM Header")
	}

	return nil
}
