package main

import (
	"fmt"
	"trainig_api/db"
	"trainig_api/model"
)

func main() {
	dbConn := db.InitDB()
	defer fmt.Printf("Successfully migrated")
	defer db.Close(dbConn)
	dbConn.AutoMigrate(&model.User{})
}
