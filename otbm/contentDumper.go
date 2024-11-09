package otbm

import "fmt"

func (otbmReader *OTBMReader) DumpFileContents() {
	fmt.Println("Dumping file contents for debugging...")

	otbmReader.resetReaderToBegginingOfTheFile()
	otbmReader.dumpContents()
}

func (otbmReader *OTBMReader) dumpContents() {
	for i := 0; i < 40; i++ {
		byteValue := make([]byte, 1)

		_, err := otbmReader.data.Read(byteValue)
		if err != nil {
			break // End of file
		}

		fmt.Printf("0x%02X ", byteValue[0])
	}

	fmt.Println("\nEnd of file dump.")
}

func (otbmReader *OTBMReader) resetReaderToBegginingOfTheFile() {
	otbmReader.data.Seek(0, 0)
}
