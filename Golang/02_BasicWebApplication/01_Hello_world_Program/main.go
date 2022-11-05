package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Fprintf takes the string and format it and then write to 'w' responsewriter
		n, err := fmt.Fprintf(w, "Hello world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written %v", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
