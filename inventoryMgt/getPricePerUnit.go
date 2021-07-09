package inventoryMgt

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	pb "vegimarket.com/serviceDefinition"
)

func GetPricePerUnit(vegiName string) *pb.GetPricePerUnitResponse {

	fname := "inventoryMgt/inventory.data"
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	inventory := &pb.Inventory{}
	// [START unmarshal_proto]
	if err := proto.Unmarshal(in, inventory); err != nil {
		log.Fatalln("Failed to parse inventory data:", err)
	}
	var price float32
	if inventory.VegiEntry[vegiName] != nil {
		price = inventory.VegiEntry[vegiName].Price
	} else {
		price = 0
	}

	return &pb.GetPricePerUnitResponse{Price: price}
}
