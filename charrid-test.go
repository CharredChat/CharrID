package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CharredChat/CharrID/CharredProcess"
)

func main() {
	var proc = CharredProcess.New(127)
	for i := 0; i < 250; i++ {
		var test, err = proc.Generate()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf(
`ID: %d
Unix Timestamp: %d
Time: %s

`, test.AsUint64(), test.GetTimestamp(), time.Unix(int64(test.GetTimestamp()/1000), 0).Format(time.RFC1123))

	}
}
