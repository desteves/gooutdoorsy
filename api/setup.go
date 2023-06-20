package api

import (
	"github.com/gin-gonic/gin"
)

const (
	RentalsEndpoint     = "/rentals"
	RentalsEndpointByID = RentalsEndpoint + "/:id"
)

func Setup(dbCnfg string) (*gin.Engine, error) {

	return gin.Default(), nil
}
