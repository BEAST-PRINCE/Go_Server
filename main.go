package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "NAMASTE!")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm Error: %v", err)
	}

	fmt.Fprintf(w, "Successful POST request\n")
	name := r.FormValue("name")
	clg := r.FormValue("college")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name: %v\nEmail: %v\nCollege: %v\n", name, email, clg)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Println("Initiating Server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
