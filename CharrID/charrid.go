package CharrID

import (
	"time"
)

// Bit lengths for ID components
const (
	LenTime = 48
	LenProc = 8
	LenIncr = 8
)

// Actual struct for IDs
type CharrID struct {
	Timestamp uint64
	Proc      uint8
	Incr      uint8
}

func New() CharrID {
	return CharrID{uint64(time.Now().UnixMilli()), 0, 0}
}

func (id *CharrID) AsUint64() uint64 {
	return uint64(id.Timestamp)<<(LenProc+LenIncr) |
		uint64(id.Proc)<<LenIncr |
		uint64(id.Incr)
}
