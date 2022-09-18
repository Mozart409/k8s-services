package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting web server...")

	message := os.Getenv("MESSAGE")
	fmt.Printf("Using message: %s\n", message)

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Request received")
		fmt.Fprintf(res, "Message: %s", message)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
