package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on port 8082")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%v", request.URL)
		_, _ = writer.Write([]byte("<h1>Welcome to gorder-v2 homepage</h1>"))
	})
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%v", request.URL)
		_, _ = writer.Write([]byte("pong"))
	})
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
