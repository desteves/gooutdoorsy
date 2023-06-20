package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

const readOneRentalSQLStatement = "SELECT * FROM rental_v WHERE id=$1"

type Postgres struct {
	Client *sql.DB
}

// queryBuilder returns the appropiate statement to send to postgres
func queryBuilder(parameters Parameters) string {
	statement := "SELECT * FROM rental_v"
	var filters []string

	if parameters.PriceMin > 0 {
		filters = append(filters, " price_per_day >= "+fmt.Sprint(parameters.PriceMin))

	}
	if parameters.PriceMax > 0 {
		filters = append(filters, " price_per_day <= "+fmt.Sprint(parameters.PriceMax))

	}
	if len(parameters.IDs) > 0 {
		// strings.Trim(strings.Replace(fmt.Sprint(parameters.IDs), " ", ",", -1), "()"
		filters = append(filters, " id IN ("+parameters.RawIDs+")")

	}

	if len(parameters.Near) == 2 {
		//  //strings.Trim(strings.Replace(fmt.Sprint(parameters.Near), " ", ",", -1), "()")
		coords := "(" + parameters.RawNear + ")"

		filters = append(filters, " "+
			"earth_box(ll_to_earth "+coords+", 100) @> ll_to_earth (lat, lng) "+
			"AND earth_distance(ll_to_earth "+coords+",  ll_to_earth (lat, lng)) < 100 ")

	}

	if len(filters) > 0 {
		statement += " WHERE " + strings.Join(filters[:], " AND ")
	}

	if parameters.Sort != "" {

		var sort string
		switch parameters.Sort {
		case "price":
			sort = "price_per_day"
		default:
			sort = parameters.Sort
		}

		statement += " ORDER BY " + sort
	}

	if parameters.Limit > 0 {
		statement += " LIMIT  " + fmt.Sprint(parameters.Limit)
	}
	if parameters.Offset > 0 {
		statement += " OFFSET " + fmt.Sprint(parameters.Offset)
	}

	fmt.Printf("%v\n", statement)
	return statement

}

func (db *Postgres) Open(conn string) (err error) {
	if conn == "" {
		return fmt.Errorf("invalid connection string")
	}
	db.Client, err = sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	return db.Client.Ping()
}

func (db *Postgres) ReadOne(id int) (resp RentalData, err error) {

	if db.Client == nil {
		return resp, fmt.Errorf("invalid Client")
	}
	err = db.Client.QueryRow(readOneRentalSQLStatement, id).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.Type,
		&resp.VehicleMake,
		&resp.VehicleModel,
		&resp.VehicleYear,
		&resp.VehicleLength,
		&resp.Sleeps,
		&resp.PrimaryImageURL,
		&resp.PricePerDay.Day,
		&resp.Location.HomeCity,
		&resp.Location.HomeState,
		&resp.Location.HomeZip,
		&resp.Location.HomeCountry,
		&resp.Location.CoordinatesLatitude,
		&resp.Location.CoordinatesLongitude,
		&resp.User.ID,
		&resp.User.FirstName,
		&resp.User.LastName)
	return resp, err
}

func (db *Postgres) ReadMany(parameters Parameters) (resp []RentalData, err error) {

	if db.Client == nil {
		return resp, fmt.Errorf("invalid Client")
	}
	var rows *sql.Rows
	rows, err = db.Client.Query(queryBuilder(parameters))
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp RentalData
		if err = rows.Scan(&tmp.ID,
			&tmp.Name,
			&tmp.Description,
			&tmp.Type,
			&tmp.VehicleMake,
			&tmp.VehicleModel,
			&tmp.VehicleYear,
			&tmp.VehicleLength,
			&tmp.Sleeps,
			&tmp.PrimaryImageURL,
			&tmp.PricePerDay.Day,
			&tmp.Location.HomeCity,
			&tmp.Location.HomeState,
			&tmp.Location.HomeZip,
			&tmp.Location.HomeCountry,
			&tmp.Location.CoordinatesLatitude,
			&tmp.Location.CoordinatesLongitude,
			&tmp.User.ID,
			&tmp.User.FirstName,
			&tmp.User.LastName); err != nil {
			return resp, fmt.Errorf("bad ReadMany Scan")
		}
		resp = append(resp, tmp)
	}

	if err = rows.Err(); err != nil {
		return resp, fmt.Errorf("bad ReadMany iteration")
	}
	return resp, err
}
