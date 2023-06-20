package database

import (
	"fmt"
)

type Mock struct {
	RentalData  []RentalData
	ReturnError bool
}

func (m *Mock) Open(conn string) (err error) {
	return nil
}

func (m *Mock) ReadOne(id int) (RentalData, error) {

	for _, rvRental := range m.RentalData {
		if rvRental.ID == id {
			return rvRental, nil
		}
	}
	return RentalData{}, fmt.Errorf("not found")

}

// TODO - add filters
func (m *Mock) ReadMany(parameters Parameters) ([]RentalData, error) {
	return m.RentalData, nil

}
