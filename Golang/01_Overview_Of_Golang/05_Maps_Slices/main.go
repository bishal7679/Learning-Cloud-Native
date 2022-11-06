package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// --------------------------- Maps ------------------------------------------------------
	myMap := make(map[string]interface{}) // first string is "Key" type and another string is "value" type and if you dont know the datatype of value then you can use interface{}
	myMap2 := make(map[string]User)

	myMap["name"] = "bishal"
	myMap["profession"] = "student"

	me := User{
		FirstName: "Rahul",
		LastName:  "Sharma",
	}

	myMap2["me"] = me

	log.Println(myMap["name"])
	log.Println(myMap["profession"])
	log.Println(myMap2["me"].FirstName, myMap2["me"].LastName)

	// --------------------------------------------- Slices ---------------------------------------------

	var mySlice []int // Slice is like an array in Golang
	mySlice2 := []string{"one", "two", "three"}

	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)
	log.Println(mySlice2)

}
