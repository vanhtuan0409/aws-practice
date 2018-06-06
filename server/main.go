package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello world"))
	})
	log.Println("Server is starting at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
