package main

import (
	"log"

	"github.com/bishal7679/gopackage/helpers"
)

func main() {
	log.Println("hello")

	var myVar helpers.Sometype // using Sometype struct of helpers package
	myVar.TypeName = "Bishal"
	myVar.TypeNumber = 12
	log.Println(myVar)

}
