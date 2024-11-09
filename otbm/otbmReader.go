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
		fmt.Println("Invalid OTBM root node detected.")
		return errors.New("invalid OTBM Magic Identifier")
	}

	return nil
}
