package main

import (
	"fmt"
	"log"

	simplepb "github.com/MickiFoerster/protobuf/go/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func jsondemo(pb proto.Message) {
	js, err := toJSON(pb)
	if err != nil {
		log.Fatalf("error while creating JSON from message: %s\n", err)
	}
	fmt.Println(js)

	sm, err := fromJSON(js)
	if err != nil {
		log.Fatalf("error while creating message from JSON: %s\n", err)
	}
	fmt.Println(sm)
}

func toJSON(pb proto.Message) (string, error) {
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(pb)
	if err != nil {
		return "", err
	}
	return out, nil
}

func fromJSON(in string) (proto.Message, error) {
	sm := &simplepb.SimpleMessage{}
	err := jsonpb.UnmarshalString(in, sm)
	if err != nil {
		return nil, err
	}
	return sm, nil
}
