package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("html")))

	if err := http.ListenAndServe(":8686", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
