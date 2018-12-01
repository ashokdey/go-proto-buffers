package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/go-proto-buffers/src/proto"
)

func main() {
	sm := doSimple()
	writeToFile("Smple.bin", sm)
}

func writeToFile(fileName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Can't serealize bytes[] to file", err)
		return err
	}

	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatal("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")
	return nil
}

func doSimple() *simplepb.Simple {
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

	fmt.Println("The message id = ", sm.GetId()) // use helper methods to stay nil safe

	return &sm
}
