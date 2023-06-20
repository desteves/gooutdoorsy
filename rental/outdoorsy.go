package rental

import (
	"github.com/desteves/gooutdoorsy/database"
)

type OutdoorsyRV struct {
	DB database.Database
}

func NewOutdoorsyProvider(conn string) (*OutdoorsyRV, error) {
	var db database.Postgres
	err := db.Open(conn)
	if err != nil {
		return nil, err
	}
	return &OutdoorsyRV{DB: &db}, nil

}

func (o *OutdoorsyRV) GetRental(id int) (database.RentalData, error) {
	return o.DB.ReadOne(id)
}

func (o *OutdoorsyRV) GetRentals(parameters database.Parameters) ([]database.RentalData, error) {
	return o.DB.ReadMany(parameters)
}
