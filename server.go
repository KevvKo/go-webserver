package main

import (
	"fmt"
	"log"
	"net/http"

)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(";8999", nil); err != nil {
		log.Fatal(err)
	}
}