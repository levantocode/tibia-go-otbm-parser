package otbm

import "encoding/binary"

func (otbmReader *OTBMReader) ReadByteAsNumber() (uint8, error) {
	var value uint8
	err := binary.Read(otbmReader.data, binary.LittleEndian, &value)

	return value, err
}

func (otbmReader *OTBMReader) Read2BytesAsNumber() (uint16, error) {
	var value uint16
	err := binary.Read(otbmReader.data, binary.LittleEndian, &value)
	return value, err
}

func (otbmReader *OTBMReader) Read4BytesAsNumber() (uint32, error) {
	var value uint32
	err := binary.Read(otbmReader.data, binary.LittleEndian, &value)
	return value, err
}

// Length-prefixed String.
func (otbmReader *OTBMReader) Read2BytesAsString() (string, error) {
	length, err := otbmReader.Read2BytesAsNumber()
	if err != nil {
		return "", err
	}

	buf := make([]byte, length)
	if _, err := otbmReader.data.Read(buf); err != nil {
		return "", err
	}

	return string(buf), nil
}
