package main

import (
	"fmt"
	"log"
	"net"
	"time"

	helloworldpb "github.com/MickiFoerster/protobuf/hello_world/protoc-helloworld"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("proto buf hello world")
	msg := &helloworldpb.HelloWorld{
		Id:           23,
		Message:      "Primzahlen",
		RandomNumber: []int32{2, 3, 5, 7, 11, 13, 17, 19},
	}
	fmt.Println(msg)
	go client(msg)
	server_done := server()
	<-server_done
}

func server() chan struct{} {
	ch := make(chan struct{})
	go func() {
		listener, err := net.Listen("tcp", ":12345")
		if err != nil {
			log.Fatalf("error: Could not create passive socket: %v\n", err)
		}
		defer listener.Close()
		fmt.Println("servers listens on ", listener.Addr())
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("error: could not accept client connection: %v\n", err)
			}
			go handleConnection(conn)
		}
		ch <- struct{}{}
	}()
	return ch
}

func handleConnection(conn net.Conn) {
	log.Println("Received client connection:")
	log.Printf("local address: %v\n", conn.LocalAddr())
	log.Printf("remote address: %v\n", conn.RemoteAddr())

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("error: could not read from connection: %v\n", err)
	}
	log.Println(n, "bytes have been read from connection")

	msg := &helloworldpb.HelloWorld{}
	err = proto.Unmarshal(buf[:n], msg)
	if err != nil {
		log.Println("unmarshal() data failed")
	}
	log.Println("unmarshaling succeeded")
	fmt.Println(msg.GetId())
	fmt.Println(msg.GetMessage())
	for _, i := range msg.GetRandomNumber() {
		fmt.Print(i, " ")
	}
	fmt.Println()

	log.Printf("closing connection ... ")
	err = conn.Close()
	if err != nil {
		log.Printf("error: could not close connection: %v\n", err)
	}
	log.Println("done")
}

func client(msg proto.Message) {
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatalf("error: could not create client connection: %v\n", err)
	}
	defer conn.Close()
	log.Println("client connection established")

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("error: could not serialize message: %v\n", err)
	}
	log.Println("protobuf message has been marshalled")

	n, err := conn.Write(data)
	if err != nil {
		log.Fatalf("error: could not write message to connection", err)
	}
	log.Println(n, "bytes have been written")
}
