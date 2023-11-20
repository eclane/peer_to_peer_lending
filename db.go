// models.go
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initialize() {
	// Open a database connection
	database, err := gorm.Open(sqlite.Open("p2p_lending.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the database
	database.AutoMigrate(&Borrower{}, &Lender{}, &Loan{})

	// Assign the global variable for database access
	db = database
}
