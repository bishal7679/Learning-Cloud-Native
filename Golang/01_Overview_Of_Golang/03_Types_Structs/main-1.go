package main

import "fmt"

var s = "seven" // variable declared outside the function scope can be accessed by all functions in that package

func main() {
	var s2 = "six"
	fmt.Println("s is ", s)
	fmt.Println("s2 is ", s2)

	saySomething("bishal")
}

func saySomething(s string) (string, string) {
	fmt.Println("s from the saySomething func is ", s)
	return s, "world"
}
