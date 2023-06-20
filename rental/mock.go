// Package rental contains the business logic. It gets called by the API and it then calls the Database.
package rental

import (
	"fmt"

	"github.com/desteves/gooutdoorsy/database"

	"golang.org/x/exp/slices"
)

type MockRV struct {
	RentalData  []database.RentalData
	ReturnError bool
}

func (m *MockRV) GetRental(id int) (database.RentalData, error) {
	for _, rvRental := range m.RentalData {
		if rvRental.ID == id {
			return rvRental, nil
		}
	}
	return database.RentalData{}, fmt.Errorf("not found")
}

func (m *MockRV) GetRentals(parameters database.Parameters) ([]database.RentalData, error) {
	var matchedRVs []database.RentalData
	for _, rvRental := range m.RentalData {
		// TODO - test other params, only doing IDs for now
		if slices.Contains(parameters.IDs, rvRental.ID) {
			matchedRVs = append(matchedRVs, rvRental)
		}
	}
	if len(matchedRVs) == 0 {
		return []database.RentalData{}, fmt.Errorf("not found")
	}
	return matchedRVs, nil
}
