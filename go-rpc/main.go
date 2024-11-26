package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func init() {

}

type server struct {
	pb.UnimplementedGreeterServer
}

// The response message containing the greetings
// type HelloReply struct
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
// }

const ClientDefaultName = "wolrd"

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("call SayHello %#v\n", in)
	return &pb.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}

var port = flag.Int("port", 50051, "server port")
var addr = flag.String("addr", "localhost:50051", "the address client connects to")
var name = flag.String("name", ClientDefaultName, "client name")

/*
Introduction: https://grpc.io/docs/what-is-grpc/introduction/
*/
func main() {

	//GOLANG: https://grpc.io/docs/languages/go/

	/**
	gRPC Server: helloworld greeter_server
	https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
	*/

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterGreeterServer(srv, &server{}) // 這一行將  request handlers 註冊到 server

	log.Printf("server listening at " + lis.Addr().String())

	go func() {
		time.Sleep(3 * time.Second)
		for {
			greeterClient()
			time.Sleep(5 * time.Second)
		}
	}()

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to server; %#v", err)
	}

}

func greeterClient() {

	time.Sleep(3 * time.Second) // waiting
	log.Printf("thread greeterClient starts\n")

	if !flag.Parsed() {
		flag.Parse()
	}

	// 1. Setup a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial failed: %#v", err)
	}
	defer conn.Close()

	// 2. Connect to server
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: *name + " at " + time.Now().GoString(),
	})

	if err != nil {
		log.Printf("error %#v", err)
	}

	log.Printf("message = %s", resp.GetMessage())

}
