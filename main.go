package main

import "net/http"

func main() {
	setupeAPI()
	http.ListenAndServe("localhost:8080", setupeAPI)
}

func setupeAPI() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}
