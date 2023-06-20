package main

import (
	"fmt"
	"os"

	"github.com/desteves/gooutdoorsy/api"
)

func main() {

	outdoorsyPostgresCnfg := os.Getenv("DATABASE_URL")
	if outdoorsyPostgresCnfg == "" {
		panic(fmt.Errorf("DATABASE_URL not set"))
	}
	port := "8080"

	api, err := api.Setup(outdoorsyPostgresCnfg)
	if err != nil {
		panic(err)
	}
	if err := api.Run("0.0.0.0:" + port); err != nil {
		panic(err)
	}

}
