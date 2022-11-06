package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Bishal",
			"last_name": "das",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Clerk",
			"last_name": "Den",
			"hair_color": "brown",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Print("Error while unmarshalling", err)

	}
	fmt.Println("Unmarshalled : ", unmarshalled)

	// write json from struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "Juli"
	m1.HairColor = "golden"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Dims"
	m2.LastName = "Juli"
	m2.HairColor = "red"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newjson, err := json.MarshalIndent(mySlice, "", "      ")
	if err != nil {
		log.Println("Error while marshalling", err)

	}

	fmt.Println("Marshalled : ", string(newjson))

}
