package main

import (
	"fmt"
	"log"

	"github.com/CharredChat/charrid/CharredProcess"
)

func main() {
	var proc = CharredProcess.New(127)
	for i := 0; i < 500; i++ {
		var test, err = proc.Generate()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Hello, World! %64b\n", test.AsUint64())

	}
}
