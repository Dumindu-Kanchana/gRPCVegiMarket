package inventoryMgt

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"math"
	"os"
	pb "vegimarket.com/serviceDefinition"
)

func AddVegetable(vegiInfo *pb.VegetableInfo) *pb.AddUpdateVegetableResponse {

	fname := "inventoryMgt/inventory.data"
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	inventory := &pb.Inventory{}
	// [START unmarshal_proto]
	if err := proto.Unmarshal(in, inventory); err != nil {
		log.Fatalln("Failed to parse inventory data:", err)
	}
	var message string
	if inventory.VegiEntry == nil {
		inventory.VegiEntry = make(map[string]*pb.VegetableInfo)
	} else if inventory.VegiEntry[vegiInfo.Name] != nil {
		message = "Vegetable already available. Please update"
		return &pb.AddUpdateVegetableResponse{Message: message}
	}  else {
		tempVegiInfo := vegiInfo
		tempVegiInfo.Amount = float32 (math.Abs(float64(vegiInfo.Amount)))
		tempVegiInfo.Price = float32 (math.Abs(float64(vegiInfo.Price)))
		inventory.VegiEntry[vegiInfo.Name] = tempVegiInfo
		message = "vegetable " + vegiInfo.Name +" added to the inventory."
	}
	// [START marshal_proto]
	out, err := proto.Marshal(inventory)
	if err != nil {
		log.Fatalln("Failed to encode the inventory data:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write inventory to file:", err)
	}
	return &pb.AddUpdateVegetableResponse{Message: message}
}
