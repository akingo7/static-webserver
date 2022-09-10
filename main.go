package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "hello" {
		http.Error( w, "error 404 not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "invalid method", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello Guys, how are you doing")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf( w, "error parsing form: %v", err)
		return
	}
	name := r.FormValue("name")
	age := r.FormValue("age")
	class := r.FormValue("class")
	fmt.Fprintf(w, "Your post is successful. Please find your details below:\nName: %v\nClass: %v\nAge: %v", name, class, age)
}



func main() {
	fileserver := http.FileServer(http.Dir("./source"))
	http.Handle( "/", fileserver )
	http.HandleFunc( "/form", formHandler )
	http.HandleFunc( "/hello", helloHandler )

	fmt.Println("Starting server on port 80...")

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}