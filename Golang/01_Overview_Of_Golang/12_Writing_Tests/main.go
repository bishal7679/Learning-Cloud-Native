package main

import (
	"errors"
	"log"
)

func main() {
	result, err := Divide(100.0, 0.0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result of division is : ", result)
}

func Divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot divide by zero")
	}

	result = x / y
	return result, nil
}
