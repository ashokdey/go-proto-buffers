package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-proto-buffers/src/proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	writeToFile("Smple.bin", sm)
	sm2 := &simplepb.Simple{}
	readFromFile("Smple.bin", sm2)
	fmt.Println(sm2)
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

func readFromFile(fileName string, pb proto.Message) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Failed to read from file", err)
		return err
	}
	err2 := proto.Unmarshal(bytes, pb)
	if err2 != nil {
		log.Fatalln("Failed to write binary to file", err2)
	}

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
