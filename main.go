package main

import (
	"fmt"
	"github.com/praveen-kumar-04/real_image_challenge/services"
)


func main(){
	locations, err := services.LoadLocations("cities.csv")

	if err != nil {
		fmt.Println("Error loading CSV:", err)
		return
	}

	distributor1 := &services.Distributor{
		Name:     "DISTRIBUTOR1",
		Includes: make(map[string]bool),
		Excludes: make(map[string]bool),
	}

	distributor1.AddInclude("INDIA")
	distributor1.AddInclude("UNITEDSTATES")
	distributor1.AddExclude("KARNATAKA-INDIA")
	distributor1.AddExclude("CHENNAI-TAMILNADU-INDIA")

	distributor2 := &services.Distributor{
		Name:     "DISTRIBUTOR2",
		Includes: make(map[string]bool),
		Excludes: make(map[string]bool),
		Parent:   distributor1,
	}

	distributor2.AddInclude("INDIA")
	distributor2.AddExclude("TAMILNADU-INDIA")


	distributor3 := &services.Distributor{
		Name:     "DISTRIBUTOR3",
		Includes: make(map[string]bool),
		Excludes: make(map[string]bool),
		Parent:   distributor2,
	}

	distributor3.AddInclude("HUBLI-KARNATAKA-INDIA")

	current_distributor := distributor1

	for location, _ := range locations {

		// fmt.Println(location, distributor1.IsallowedtoDistribute(location,locations[location]))

		if current_distributor.IsallowedtoDistribute(location,locations[location]) {
			fmt.Printf("%s is allowed to distribute in %s-%s-%s :=> Yes\n",current_distributor.Name,location,locations[location][0],locations[location][1])
			
		}else{
			fmt.Printf("%s is allowed to distribute in %s-%s-%s :=> No\n",current_distributor.Name,location,locations[location][0],locations[location][1])
		}
	}

}