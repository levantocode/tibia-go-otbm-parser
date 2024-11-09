package otbm

import "fmt"

func (otbmReader *OTBMReader) ReadHeader() (bool, error) {
	version, err := otbmReader.Read4BytesAsNumber()
	if err != nil {
		return false, err
	}
	fmt.Println("Header - Version:", version)

	mapWidth, err := otbmReader.Read2BytesAsNumber()
	if err != nil {
		return false, err
	}
	fmt.Println("Header - Map Width:", mapWidth)

	mapHeight, err := otbmReader.Read2BytesAsNumber()
	if err != nil {
		return false, err
	}
	fmt.Println("Header - Map Height:", mapHeight)

	itemVersionMajor, err := otbmReader.Read4BytesAsNumber()
	if err != nil {
		return false, err
	}
	fmt.Println("Header - Item Major Version:", itemVersionMajor)

	itemVersionMinor, err := otbmReader.Read4BytesAsNumber()
	if err != nil {
		return false, err
	}
	fmt.Println("Header - Item Minor Version:", itemVersionMinor)

	otbmReader.DumpFileContents()

	return true, nil
}
