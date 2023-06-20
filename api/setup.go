package api

import (
	"github.com/desteves/gooutdoorsy/rental"
	"github.com/gin-gonic/gin"
)

const (
	RentalsEndpoint     = "/rentals"
	RentalsEndpointByID = RentalsEndpoint + "/:id"
)

func Setup(dbCnfg string) (*gin.Engine, error) {
	outdoorsy, err := rental.NewOutdoorsyProvider(dbCnfg)
	if err != nil {
		return nil, err
	}
	api := API{
		RVRentalProvider: outdoorsy,
	}

	webRouter := gin.Default()
	// healthcheck
	webRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	webRouter.GET(RentalsEndpoint, api.GetRVRentalsHandler)
	webRouter.GET(RentalsEndpointByID, api.GetRVRentalByIDHandler)

	return webRouter, nil
}
