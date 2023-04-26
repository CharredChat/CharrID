package main

import (
	"fmt"
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

func (id *CharrID) as_u64() uint64 {
	return uint64(id.Timestamp)<<(LenProc+LenIncr) |
		uint64(id.Proc)<<LenIncr |
		uint64(id.Incr)
}

func main() {
	var test = CharrID{68624, 23, 96}
	fmt.Printf("Hello, World! %64b", test.as_u64())
}
