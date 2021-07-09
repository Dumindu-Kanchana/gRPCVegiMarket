package inventoryMgt

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	pb "vegimarket.com/serviceDefinition"
)

func GetVegiNames() *pb.GetVegetableNamesResponse {

	fname := "inventoryMgt/inventory.data"
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	inventory := &pb.Inventory{}

	if err := proto.Unmarshal(in, inventory); err != nil {
		log.Fatalln("Failed to parse inventory data:", err)
	}

	var vegiNamesResponse = &pb.GetVegetableNamesResponse{}
	vegiNames := make([]string, len(inventory.VegiEntry))
	i := 0
	for key := range inventory.VegiEntry {
		vegiNames[i] = key
		i++
	}
	vegiNamesResponse.Name = vegiNames
	return vegiNamesResponse
}

