package main

import (
	"log"
	"net/http"
	"time"
)

func pong(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	w.Write([]byte("pong"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pong)

	server := &http.Server{
		Addr: 				":8080",
		WriteTimeout: time.Second * 5,
		Handler: 			mux,
	}

	log.Println("Starting server at localhost:8080")
	log.Fatal(server.ListenAndServe())
}