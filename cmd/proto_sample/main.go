package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/wgarcia4190/protobuf_sample_go/internal/domain/messages"
)

func main() {
	sm := doSimple()
	readAndWriteToFile(sm)
	json(sm)
}

func readAndWriteToFile(sm proto.Message) {


	if err := writeToFile("simple.bin", sm); err != nil {
		panic(err)
	}

	sm2 := &simplepb.SimpleMessage{}

	if err := readFromFile("simple.bin", sm2); err != nil {
		panic(err)
	}

	fmt.Println(sm2)
}

func json(sm proto.Message){

	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)

	fmt.Println(sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Could not unmarshal", err)
	}
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0600); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Could not put the bytes into the proto struct", err)
		return err
	}

	fmt.Println("Data has been read!")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id: 12345,
		IsAvailable: true,
		Name: "Simple Name",
		SimpleList: []int32{1, 2, 3, 4, 5},
	}

	return &sm
}
