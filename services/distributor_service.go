package services

import(
	"strings"
)

type Distributor struct {
	Name      string
	Includes  map[string]bool
	Excludes  map[string]bool
	Parent    *Distributor
}

func (d *Distributor) AddInclude(location string) {
	location = strings.ToUpper(location)
	d.Includes[location] = true
}

func (d *Distributor) AddExclude(location string) {
	location = strings.ToUpper(location)
	d.Excludes[location] = true
}


func (d *Distributor) IsallowedtoDistribute(location string, locations []string) bool {
	
	// checking the cities, state, country in the exluded list to identify if the location is allowed to distribute
	full :=strings.Join([]string{locations[0], locations[1]}, "-")
	if d.Excludes[locations[1]] || d.Excludes[strings.Join([]string{locations[0], locations[1]}, "-")] || d.Excludes[strings.Join([]string{full,location}, "-")] {
		return false
	}


	// checking the cities, state, country in the included list
	if d.Includes[locations[1]] || d.Includes[strings.Join([]string{locations[0], locations[1]}, "-")] || d.Includes[strings.Join([]string{full,location}, "-")] {
		if d.Parent != nil {
			return d.Parent.IsallowedtoDistribute(location, locations)
		}
		return true
	}

	if d.Parent != nil {
		return d.Parent.IsallowedtoDistribute(location, locations)
	}

	return false
}