package main

import (
	"log"

	"github.com/lucasbarroso23/ecomm-mysql/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected to database")
}
