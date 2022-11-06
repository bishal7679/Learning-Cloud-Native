package main

import "log"

type Animal interface {
	says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "loki",
		Breed: "German Shepard",
	}

	PrintInfo(dog)
	gorilla := Gorilla{
		Name:          "Kin Kong",
		Color:         "black",
		NumberOfTeeth: 32,
	}
	PrintInfo(gorilla)
}

func (d Dog) says() string {
	return "woof"
}
func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) says() string {
	return "woof"
}
func (g Gorilla) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.says(), "and has", a.NumberOfLegs(), "legs")
}
