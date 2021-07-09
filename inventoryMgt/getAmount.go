package inventoryMgt

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	pb "vegimarket.com/serviceDefinition"
)

func GetAmount(vegiName string) *pb.GetAmountResponse {

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
	var amount float32
	if inventory.VegiEntry[vegiName] != nil {
		amount = inventory.VegiEntry[vegiName].Amount
	} else {
		amount = 0
	}

	return &pb.GetAmountResponse{Amount: amount}
}
