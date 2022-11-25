package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/w-woong/product/dto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	proto.UnimplementedProductServer
}

func (s *server) ReadProduct(ctx context.Context, in *proto.ProductRequest) (*proto.ProductResponse, error) {
	return &proto.ProductResponse{Id: in.Id, Name: "haha"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":59001")
	if err != nil {
		fmt.Println(err)
		return
	}

	svr := grpc.NewServer()
	proto.RegisterProductServer(svr, &server{})

	go func() {
		time.Sleep(2 * time.Second)
		request()
	}()
	svr.Serve(lis)
}

func request() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":59001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewProductClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ReadProduct(ctx, &proto.ProductRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())
}
