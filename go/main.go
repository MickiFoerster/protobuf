package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/MickiFoerster/protobuf/go/src/complex"
	simplepb "github.com/MickiFoerster/protobuf/go/src/simple"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	const fname = "simple.bin"
	err := writeToFile(fname, sm)
	if err != nil {
		log.Fatalf("error while writing message: %s\n", err)
	}

	pb, err := readFromFile(fname)
	if err != nil {
		log.Fatalf("error while reading message from file: %s\n", err)
	}
	fmt.Println(pb)

	jsondemo(pb)

	doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "first message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "third message",
			},
		},
	}
	fmt.Println(cm)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		return err
	}

	fmt.Println("Data has been written")
	return nil
}

func readFromFile(fname string) (proto.Message, error) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	sm := simplepb.SimpleMessage{}
	err = proto.Unmarshal(in, &sm)
	if err != nil {
		return nil, err
	}
	return &sm, nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{2, 3, 5, 7, 11, 13},
	}

	fmt.Println(sm)

	sm.Name = "New name"
	fmt.Println(sm, " -> ", sm.GetId())
	return &sm
}
