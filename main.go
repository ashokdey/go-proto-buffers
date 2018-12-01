package main

import (
	"fmt"

	"github.com/go-proto-buffers/src/proto"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.Simple{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)

	// we can manipulate the attributes
	sm.Name = "Renamed my message"
	fmt.Println(sm)

	fmt.Println("The message id = ", sm.GetId())
}
