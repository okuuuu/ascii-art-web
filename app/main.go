package main

import (
	"log"
	"main/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.MakeHandler(server.HomeHandler))
	http.HandleFunc("/ascii-art", server.MakeHandler(server.GenerateHandler))

	log.Print("Starting server at port 8080...\n")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Print("http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
