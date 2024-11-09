package otbm

import "fmt"

func (r *OTBMReader) hasValidMagicIdentifier() (bool, error) {
	magicIdentifier, err := r.Read4BytesAsNumber() // Should be a 4 Byte String but works for empty
	if err != nil {
		return false, err
	}

	if isInvalid(magicIdentifier) {
		return false, fmt.Errorf(
			"unexpected node type: got %d, expected %d",
			magicIdentifier, OTBM_INTRO_LITERAL,
		)
	}

	return true, nil
}

func isInvalid(magicIdentifier uint32) bool {
	return magicIdentifier != OTBM_INTRO_LITERAL &&
		magicIdentifier != OTBM_INTRO_EMPTY
}
