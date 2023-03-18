package main

import (
	"log"
	"net/http"
)

func main() {
	setupeAPI()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupeAPI() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}
