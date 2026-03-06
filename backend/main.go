package main

import (
	"log"

	"github.com/RemyPaulJr/portfolio-go/backend/config"
)

func main() {

	dbConnection := config.EstablishConnection()

	log.Print("Successful db connection:", dbConnection)
}
