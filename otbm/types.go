package otbm

import "bytes"

// Node types
const (
	OTBM_INTRO_LITERAL       = 0x4D42544F // ("OTBM")
	OTBM_INTRO_EMPTY         = 0x00000000
	OTBM_ROOT                = 0x01
	OTBM_MAP_DATA            = 0x02
	OTBM_TILE_AREA           = 0x03
	OTBM_TILE                = 0x04
	OTBM_HOUSETILE           = 0x05
	OTBM_TOWNS               = 0x06
	OTBM_TOWN                = 0x07
	OTBM_WAYPOINTS           = 0x08
	OTBM_WAYPOINT            = 0x09
	OTBM_ATTR_DESCRIPTION    = 0x01
	OTBM_ATTR_TILE_FLAGS     = 0x0B
	OTBM_ATTR_EXT_HOUSE_FILE = 0x13
)

// Reads & Processes an OTBM file.
type OTBMReader struct {
	data *bytes.Reader
}

// Represents a Map Coordinate.
type Position struct {
	X, Y, Z uint16
}
