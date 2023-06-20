package rental

import (
	"github.com/desteves/gooutdoorsy/database"
)

type RV interface {
	GetRental(id int) (database.RentalData, error)
	GetRentals(parameters database.Parameters) ([]database.RentalData, error)
}
