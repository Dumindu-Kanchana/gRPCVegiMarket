package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
	pb "vegimarket.com/serviceDefinition"
)

const (
	address = "localhost:50050"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s Vegitable Name is required!\n", os.Args[0])
	}
	vegiName := os.Args[1]

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVegiMarketClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.GetPricePerUnit(ctx, &pb.GetPricePerUnitRequest{Name: vegiName})
	if err != nil {
		log.Fatalf("could not obtain the vegitable price: %v", err)
	}
	fmt.Println(response.Price)
}
