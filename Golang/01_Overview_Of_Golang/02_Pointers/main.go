package main

import "fmt"

func main() {
	var myString string
	myString = "Green"

	fmt.Println("mystring is set to ", myString)
	changeUsingPointer(&myString)
	fmt.Println("after func call mystring is set to ", myString)
}

func changeUsingPointer(s *string) {
	fmt.Println("s is set to this address : ", s)
	newValue := "red"
	*s = newValue // changing the value "green" to "red" in original address of myString
}
