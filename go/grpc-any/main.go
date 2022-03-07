package main

//
// https://pkg.go.dev/google.golang.org/protobuf@v1.27.1/types/known/anypb
//

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/anypb"

	"github.com/egustafson/sandbox/go/grpc-any/pb"
)

func main() {

	sentMessage := "field-value-in-str-message-in-any-message"
	anyMsg := marshalAnyMessage(sentMessage)

	receivedMessage := unmarshalAnyMessage(anyMsg)

	log.Printf("sent:      %s\n", sentMessage)
	log.Printf("received:  %s\n", receivedMessage)

	fmt.Println("done.")
}

func marshalAnyMessage(msg string) *pb.AnyMessage {

	strMsg := &pb.StrMessage{Msg: msg}

	any, err := anypb.New(strMsg) // <-- use anypb.New( concrete-type ) to create an 'any'  ***
	if err != nil {
		log.Fatalf("failed to convert strMsg to any: %v", err)
	}

	anyMsg := &pb.AnyMessage{
		Anything: any,
	}

	return anyMsg
}

func unmarshalAnyMessage(anyMsg *pb.AnyMessage) string {

	m, err := anyMsg.Anything.UnmarshalNew()
	if err != nil {
		log.Fatalf("failed to unmarshal any: %v", err)
	}

	switch m := m.(type) { // <-- use type switch to detect actual type of the field  ***
	case *pb.StrMessage:
		return m.Msg
	}

	return "--* unknown type returned *--"
}
