package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World\n")
	})

	fmt.Println("Server started on port 9091")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		panic(err)
	}
}
