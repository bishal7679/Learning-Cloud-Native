package main

import (
	"log"
)

type User struct {
	FirstName  string
	LastName   string
	profession string
}

func main() {
	for i := 0; i <= 10; i++ { // normal for loop
		log.Println(i)
	}

	// -------------------for loop over slices-------------------------------
	mySlice := []string{"java", "C++", "Golang", "sql", "sass"}

	for _, x := range mySlice {
		log.Println(x)
	}

	// ---------------------for loop over maps---------------------------------
	myMap := make(map[string]string)
	myMap["cat"] = "cat"
	myMap["dog"] = "dog"
	myMap["deer"] = "deer"
	for i, x := range myMap {
		log.Println(i, x)
	}

	// -------------------------for loop over struct---------------------------
	var mySlice2 []User

	user1 := User{
		FirstName:  "bishal",
		LastName:   "das",
		profession: "Engineer",
	}

	user2 := User{
		FirstName:  "John",
		LastName:   "Doe",
		profession: "Doctor",
	}

	mySlice2 = append(mySlice2, user1)
	mySlice2 = append(mySlice2, user2)

	for _, x := range mySlice2 {
		log.Println(x.FirstName, x.LastName)
	}

}
