package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	route2 "gitlab.com/Tonybsilva-dev/code-delivery-simulator/application/route"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
