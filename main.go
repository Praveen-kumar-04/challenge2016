package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/praveen-kumar-04/real_image_challenge/services"
)

func checking_list_cities(current_distributor *services.Distributor,locations map[string][]string){
	for location, _ := range locations {
		if current_distributor.IsallowedtoDistribute(location,locations[location]) {
			fmt.Printf("%s is allowed to distribute in %s-%s-%s :=> Yes\n",current_distributor.Name,location,locations[location][0],locations[location][1])
			
		}else{
			fmt.Printf("%s is allowed to distribute in %s-%s-%s :=> No\n",current_distributor.Name,location,locations[location][0],locations[location][1])
		}
	}	
	
}

func distributors_list(list_of_distributors []*services.Distributor){
	fmt.Println("List of Distributors")
		for ind,distributor := range list_of_distributors {
			fmt.Println(ind+1,"=>",distributor.Name)
		}
}

func read_data_from_user() int{
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the ID of the distributor")
		distributor_id, _ := reader.ReadString('\n')
		distributor_id = strings.TrimSuffix(distributor_id, "\n")
		distributor_id = strings.TrimSuffix(distributor_id, "\r")
		id, _ := strconv.Atoi(distributor_id)
		return id

}


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


	list_of_distributors := []*services.Distributor{distributor1,distributor2,distributor3}

	for {
		distributors_list(list_of_distributors)
		checking_list_cities(list_of_distributors[read_data_from_user()-1],locations)
		fmt.Println("Do you want to check for another distributor? (yes/no)")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSuffix(choice, "\n")
		choice = strings.TrimSuffix(choice, "\r")
		if choice == "no" {	
			break
		}
}
}