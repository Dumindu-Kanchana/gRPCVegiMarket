package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
	"strconv"
	pb "vegimarket.com/serviceDefinition"
)

const (
	address = "localhost:50050"
)

func main() {

	if len(os.Args) != 4 {
		log.Fatalf("Usage:  %s Vegitable information required!\n", os.Args[0])
	}
	vegiName := os.Args[1]
	price, err := strconv.ParseFloat(os.Args[2], 32)
	amount, err := strconv.ParseFloat(os.Args[3], 32)
	
	if err != nil {
		fmt.Printf("Error while parsing the arguments")
		os.Exit(93)
	}

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
	response, err := c.AddVegetable(ctx, &pb.VegetableInfo{Name: vegiName, Amount: float32(amount), Price: float32(price)})
	if err != nil {
		log.Fatalf("could not obtain add vegitable response: %v", err)
	}
	fmt.Println(response.Message)
}
