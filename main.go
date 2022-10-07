package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	w.Write([]byte("Welcome!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Viewing a snippet " + r.URL.Query().Get("id")))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating snippet!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippets/view", snippetView)
	mux.HandleFunc("/snippets/create", snippetCreate)

	log.Println("Starting application on server with port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
