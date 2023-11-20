// models.go
package main

import "gorm.io/gorm"

type Borrower struct {
	gorm.Model
	Name  string
	Email string
}

type Lender struct {
	gorm.Model
	Name  string
	Email string
}

type Loan struct {
	gorm.Model
	BorrowerID uint
	LenderID   uint
	Amount     float64
	Status     string // Pending, Approved, Rejected, Completed, etc.
}
