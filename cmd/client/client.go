package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close() // Verify when this attribute are in idle

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	for {
		stream, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive data from gRPC Server: %v", err)
		}

		fmt.Println(stream)

	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "id1",
			Name:  "Eduardo",
			Email: "e@e.com",
		},
		&pb.User{
			Id:    "id2",
			Name:  "Marcelo",
			Email: "m@m.com",
		},
		&pb.User{
			Id:    "id3",
			Name:  "Sapo",
			Email: "s@s.com",
		},
		&pb.User{
			Id:    "id4",
			Name:  "Zim",
			Email: "z@z.com",
		},
		&pb.User{
			Id:    "id5",
			Name:  "Gugu",
			Email: "g@g.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating conn: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receive data: %v", err)
	}

	fmt.Println(res)

}
