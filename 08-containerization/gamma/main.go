package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GAMMA")
	})

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8082", h))
}
