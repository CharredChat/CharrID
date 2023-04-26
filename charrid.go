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

func main() {
	fmt.Println("Hello, World! %s", CharrID{18446744073709551615, 64, 40})
}
