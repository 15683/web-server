package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// func viewHandler(writer http.ResponseWriter, request *http.Request) {
// 	message := []byte("Hello, web!")
// 	_, err := writer.Write(message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

type Person struct {
	Name string
	Age  int
}

var people []Person

func main() {
	// http.HandleFunc("/hello", viewHandler)
	// err := http.ListenAndServe("localhost:8080", nil)
	// log.Fatal(err)

	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/health", healthCheckHandler)

	log.Println("server start listening on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w)
	case http.MethodPost:
		postPerson(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(people)
	fmt.Fprintf(w, "get people: '%v'", people)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	people = append(people, person)
	fmt.Fprintf(w, "post new person: '%v'", person)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "http web-server works correctly")
}
