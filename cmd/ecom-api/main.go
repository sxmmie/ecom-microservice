package ecomapi

import (
	"log"

	"github.com/sxmmie/ecom-microservice/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	defer db.Close()
	log.Println("successfully connected to the database")

	// Define store mothod and pass in db as args
}
