package main

import (
	"fmt"

	CharrID "github.com/CharredChat/charrid/CharrID"
)

func main() {
	var test = CharrID.New()
	fmt.Printf("Hello, World! %64b", test.AsUint64())
}
