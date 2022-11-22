package main

import (
	"fmt"
	"log"
	"net/http"
)

var listOfThanks []string

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/getAllThanks", returnAllThanks)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func returnAllThanks(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getAllThanks" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "All Thanks: %s\n", listOfThanks)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	thankfulFor := r.FormValue("thankfulFor")

	listOfThanks = append(listOfThanks, thankfulFor)

	fmt.Fprintf(w, "\n You are %s\n", name)
	fmt.Fprintf(w, "You're thankful for %s\n", thankfulFor)
	fmt.Fprintf(w, "Thank you for submiting!")
}
