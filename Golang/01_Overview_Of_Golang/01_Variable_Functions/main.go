package main

import "fmt"

func main() {

	// Declaring a variable
	var something string = "Hello world!"
	var num int = 5
	// Printing the variable
	fmt.Println(something)
	fmt.Println(num)

	// Calling a function
	some1, some2 := Somefunc("Goodbye!")
	fmt.Println(some1 + " " + some2)

}

func Somefunc(str string) (string, string) {
	return str, "bishal"
}
