package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"log"
	"os"
	pb "vegimarket.com/serviceDefinition"
)

func main() {

	fname := "inventoryMgt/inventory.data"
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	inventory := &pb.Inventory{}

	if err := proto.Unmarshal(in, inventory); err != nil {
		log.Fatalln("Failed to parse inventory data:", err)
	}
	// [END unmarshal_proto]

	listInventory(os.Stdout, inventory)

}

func listInventory(w io.Writer, inventory *pb.Inventory) {
	for _, p := range inventory.VegiEntry {
		printEntries(w, p)
	}
}

func printEntries(w io.Writer, p *pb.VegetableInfo) {
	fmt.Fprintln(w, "Name:", p.Name)
	fmt.Fprintln(w, "Price:", p.Price)
	fmt.Fprintln(w, "Amount:", p.Amount)
	fmt.Println("##################")
}
