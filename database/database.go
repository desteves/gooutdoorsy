// Package database performs CRUD operations against business logic data
package database

type Database interface {
	Open(conn string) (err error)

	ReadOne(id int) (RentalData, error)
	ReadMany(parameters Parameters) ([]RentalData, error)
}

type RentalData struct {
	ID              int          `json:"id" db:"id"`
	Name            string       `json:"name" db:"name"`
	Description     string       `json:"description" db:"description"`
	Type            string       `json:"type" db:"type"`
	VehicleMake     string       `json:"make" db:"vehicle_make"`
	VehicleModel    string       `json:"model" db:"vehicle_model"`
	VehicleYear     int          `json:"year" db:"vehicle_year"`
	VehicleLength   float64      `json:"length" db:"vehicle_length"`
	Sleeps          int          `json:"sleeps" db:"sleeps"`
	PrimaryImageURL string       `json:"primary_image_url" db:"primary_image_url"`
	PricePerDay     PriceData    `json:"price_per_day"`
	Location        LocationData `json:"location"`
	User            UserData     `json:"user"`
}

type PriceData struct {
	Day int `json:"day" db:"price_per_day"`
}

type LocationData struct {
	HomeCity             string  `json:"city" db:"home_city"`
	HomeState            string  `json:"state" db:"home_state"`
	HomeZip              string  `json:"zip" db:"home_zip"`
	HomeCountry          string  `json:"country" db:"home_country"`
	CoordinatesLatitude  float64 `json:"lat" db:"lat"`
	CoordinatesLongitude float64 `json:"lng" db:"lng"`
}

type UserData struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

type Parameters struct {
	PriceMin int    `form:"price_min"`
	PriceMax int    `form:"price_max"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
	Sort     string `form:"sort"`
	RawIDs   string `form:"ids"`
	RawNear  string `form:"near"`
	IDs      []int
	Near     []float64
}
