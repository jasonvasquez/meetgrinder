package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	fmt.Println("About to Listen")
	log.Fatal(http.ListenAndServe(":8081", router))
	fmt.Println("Listening!")
}
