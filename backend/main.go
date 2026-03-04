package main

import (
	"fmt"
	"log"

	"github.com/RemyPaulJr/portfolio-go/backend/config"
)

func main() {
	fmt.Print("Hello, world!")

	dbConnection := config.EstablishConnection()

	log.Print(dbConnection)
}
