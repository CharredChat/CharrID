package CharrID

import (
	"errors"
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
	timestamp uint64
	proc      uint8
	incr      uint8
}

func New(proc uint8, incr uint8) (CharrID, error) {
	var time = uint64(time.Now().UnixMilli())
	if time >= 1<<LenTime {
		return CharrID{}, errors.New("Hello future man! unfortunately, this software is not programmed to handle the current date. Sorry.")
	}
	return CharrID{time, proc, incr}, nil
}

func (id *CharrID) AsUint64() uint64 {
	return uint64(id.timestamp)<<(LenProc+LenIncr) |
		uint64(id.proc)<<LenIncr |
		uint64(id.incr)
}

func (id *CharrID) GetTimestamp() uint64 {
	return id.timestamp
}
