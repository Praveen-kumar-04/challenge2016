package services
import (
	"encoding/csv"
	"os"
	"strings"
)

// Loading the csv file and storing the city, state, country in the map
func LoadLocations(fileloc string) (map[string][]string, error) {
	locations := make(map[string][]string)
	file, err := os.Open(fileloc)
	if err != nil {
		return locations, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return locations, err
	}

	for _, line := range lines[1:] {
		city, province, country :=  line[3], line[4], line[5]
		locations[strings.ToUpper(city)] = []string{strings.ToUpper(province),strings.ToUpper(country)}
	}

	return locations, nil
}