package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
	pb "vegimarket.com/serviceDefinition"
)

func main() {
	fname := "inventoryMgt/inventory.data"
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}
	// [START unmarshal_proto]
	inventory := &pb.Inventory{}
	if err := proto.Unmarshal(in, inventory); err != nil {
		log.Fatalln("Failed to parse inventory data:", err)
	}
	inventory.VegiEntry = make(map[string]*pb.VegetableInfo)
	// [START marshal_proto]
	out, err := proto.Marshal(inventory)
	if err != nil {
		log.Fatalln("Failed to encode the inventory data:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write inventory to file:", err)
	}
}
