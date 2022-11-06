package main

import "fmt"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string { // This (m *myStruct) is called "Receivers"
	return m.FirstName
}
func main() {
	var myVar myStruct
	myVar.FirstName = "Bishal"

	myVar2 := myStruct{
		FirstName: "Rahul",
	}

	// fmt.Println("myVar is set to", myVar.FirstName)
	// fmt.Println("myVar2 is set to", myVar2.FirstName)

	fmt.Println("myVar is set to", myVar.printFirstName())   // with myStruct the func printFirstName is associated so we can use this func under myStruct
	fmt.Println("myVar2 is set to", myVar2.printFirstName())

}
