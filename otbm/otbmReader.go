package otbm

import (
	"bytes"
	"os"
)

// Reads & Processes an OTBM file.
type OTBMReader struct {
	data *bytes.Reader
}

// Initializes a new OTBMReader.
func NewOTBMReader(filePath string) (*OTBMReader, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &OTBMReader{data: bytes.NewReader(file)}, nil
}
