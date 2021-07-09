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

	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s Vegitable Name is required!\n", os.Args[0])
	}
	vegiName := os.Args[1]
	fname := "inventoryMgt/inventory.data"
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	// [START marshal_proto]
	inventory := &pb.Inventory{}
	// [START_EXCLUDE]
	if err := proto.Unmarshal(in, inventory); err != nil {
		log.Fatalln("Failed to parse inventory data:", err)
	}

	delete(inventory.VegiEntry, vegiName)

	out, err := proto.Marshal(inventory)
	if err != nil {
		log.Fatalln("Failed to encode the inventory data:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write inventory to file:", err)
	}
}
