# Real Image Challenge 2016

This project implements a hierarchical distribution rights management system in Go.

## Code Structure

### Distributor Structure
```go
type Distributor struct {
    Name      string
    Includes  map[string]bool
    Excludes  map[string]bool
    Parent    *Distributor
}
```
The Distributor struct represents a distribution entity with:
- Name: Identifier for the distributor
- Includes: Locations where distribution is allowed
- Excludes: Locations where distribution is blocked
- Parent: Reference to parent distributor for hierarchical rights

### Key Functions

1. **LoadLocations**
   - Reads location data from a CSV file
   - Creates a map of cities with their corresponding province and country
   - Format: `map[CITY] = [PROVINCE, COUNTRY]`

2. **IsallowedtoDistribute**
   - Determines if distribution is allowed for a given location
   - Checks hierarchy in following order:
     1. Excluded locations (city, state-country, city-state-country)
     2. Included locations (city, state-country, city-state-country)
     3. Parent distributor permissions

### Distribution Rules
- Excludes take precedence over includes
- If location is not explicitly included/excluded, check parent distributor
- If no parent exists and no explicit inclusion, distribution is denied

## Example Usage

```go
// Create root distributor
distributor1 := &Distributor{
    Name:     "DISTRIBUTOR1",
    Includes: make(map[string]bool),
    Excludes: make(map[string]bool),
}

// Set permissions
distributor1.AddInclude("INDIA")
distributor1.AddExclude("KARNATAKA-INDIA")

// Check distribution rights
isAllowed := distributor1.IsallowedtoDistribute(location, locations[location])
```

## CSV File Format
The program expects a CSV file with the following columns:
- Column 4: City
- Column 5: Province/State
- Column 6: Country

## Dependencies
- Standard Go libraries:
  - encoding/csv
  - strings
  - fmt
  - os

## Running the Program
```bash
go run main.go
```
The program will read from `cities.csv` and output distribution rights for each location.

