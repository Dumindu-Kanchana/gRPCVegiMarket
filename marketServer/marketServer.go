package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	im "vegimarket.com/inventoryMgt"
	pb "vegimarket.com/serviceDefinition"
)

const (
	port = "50050"
)

type server struct {
	pb.UnimplementedVegiMarketServer
}

func (s *server) GetVegetableNames(ctx context.Context, in *pb.GetVegetableNamesRequest) (*pb.GetVegetableNamesResponse, error) {
	log.Printf("GetVegitableNames request Received")
	return im.GetVegiNames(), nil
}

func (s *server) GetPricePerUnit(ctx context.Context, in *pb.GetPricePerUnitRequest) (*pb.GetPricePerUnitResponse, error) {
	log.Printf("GetPricePerUnit request Received")
	return im.GetPricePerUnit(in.Name), nil
}

func (s *server) GetAmount(ctx context.Context, in *pb.GetAmountRequest) (*pb.GetAmountResponse, error) {
	log.Printf("GetAmount request Received")
	return im.GetAmount(in.Name), nil
}

func (s *server) AddVegetable(ctx context.Context, in *pb.VegetableInfo) (*pb.AddUpdateVegetableResponse, error) {
	log.Printf("AddVegetable request Received")
	return im.AddVegetable(in), nil
}

func (s *server) UpdateVegetable(ctx context.Context, in *pb.VegetableInfo) (*pb.AddUpdateVegetableResponse, error) {
	log.Printf("UpdateVegetable request Received")
	return im.UpdateVegetable(in), nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVegiMarketServer(s,&server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
