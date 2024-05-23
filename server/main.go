package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/endpoint", endpoint)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fmt.Println(time.Now())
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func endpoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Endpoint :)")
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
