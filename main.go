// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("P2P Lending Platform")

	// Initialize database, router, and other components
	initialize()

	// Define routes
	http.HandleFunc("/borrowers", handleBorrowers)
	http.HandleFunc("/lenders", handleLenders)
	http.HandleFunc("/loans", handleLoans)

	// Start server
	http.ListenAndServe(":8080", nil)
}
