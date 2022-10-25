package main

import (
	"fmt"
	"net/http"
	"os"
)

func appHandler(w http.ResponseWriter, req *http.Request) {
	response := "this is a test"
	response = "bad-override"

	fmt.Fprintf(w, "%s\n", response)
}

func main() {
	http.HandleFunc("/app", appHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("error serving webserver: %s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
