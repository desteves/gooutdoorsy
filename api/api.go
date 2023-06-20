// Package api has a business logic provider and all the handlers for such
package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/desteves/gooutdoorsy/database"
	"github.com/desteves/gooutdoorsy/rental"

	"github.com/gin-gonic/gin"
)

type API struct {
	RVRentalProvider rental.RV
}

const (
	ErrorResponseField                = "message"
	ErrorMsgStatusBadRequest          = "invalid request"
	ErrorMsgStatusNotFound            = "not found"
	ErrorMsgStatusInternalServerError = "something went wrong"
)

// GetRVRentalByIDHandler returns at most one rv rental that matches the given id
func (a *API) GetRVRentalByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ErrorResponseField: ErrorMsgStatusBadRequest})
	}
	bodyResponse, err := a.RVRentalProvider.GetRental(id)
	if err != nil {
		if err.Error() == ErrorMsgStatusNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ErrorResponseField: ErrorMsgStatusNotFound})
		}
		// TODO -- check other errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ErrorResponseField: ErrorMsgStatusInternalServerError})
	}
	c.JSON(http.StatusOK, bodyResponse)

}

func bindQueryHelper(parameters *database.Parameters) error {
	if parameters.RawIDs != "" {
		for _, stringID := range strings.Split(parameters.RawIDs, ",") {
			id, err := strconv.Atoi(stringID)
			if err != nil {
				return err
			}
			parameters.IDs = append(parameters.IDs, id)
		}
	}

	if parameters.RawNear != "" {
		for _, stringLoc := range strings.Split(parameters.RawNear, ",") {
			loc, err := strconv.ParseFloat(stringLoc, 64)
			if err != nil {
				return err
			}
			parameters.Near = append(parameters.Near, loc)
		}

		if len(parameters.Near) != 2 {
			return fmt.Errorf(ErrorMsgStatusBadRequest)
		}

	}
	return nil

}

// GetRVRentalsHandler  fetches any rentals that match the query parameters
func (a *API) GetRVRentalsHandler(c *gin.Context) {
	var parameters database.Parameters

	err := c.BindQuery(&parameters)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ErrorResponseField: ErrorMsgStatusBadRequest})
	}
	err = bindQueryHelper(&parameters)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ErrorResponseField: ErrorMsgStatusBadRequest})
	}

	bodyResponse, err := a.RVRentalProvider.GetRentals(parameters)
	if err != nil {
		if err.Error() == ErrorMsgStatusNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ErrorResponseField: ErrorMsgStatusNotFound})
		}
		// TODO -- check other errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ErrorResponseField: ErrorMsgStatusInternalServerError})
	}
	c.JSON(http.StatusOK, bodyResponse)
}
