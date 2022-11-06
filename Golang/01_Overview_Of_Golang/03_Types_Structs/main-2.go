package main

import (
	"fmt"
	"time"
)

// This is a struct type of User
type User struct {
	// Taking the first letter capital of variable/function we can access that variable/function outside the package
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Bishal",
		LastName:    "das",
		PhoneNumber: "1234567890",
	}

	fmt.Println(user.FirstName, user.LastName, user.PhoneNumber, user.BirthDate) // BirthDate is not taken in user variable so it will throw default value
}
