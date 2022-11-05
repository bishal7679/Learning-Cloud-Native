package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, err := AddValue(2, 2)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

func AddValue(x, y int) (int, error) {
	sum := x + y
	return sum, nil
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
