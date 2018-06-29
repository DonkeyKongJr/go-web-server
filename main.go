package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
}

func provideTestPeople(){
	people = append(people, Person{ID: "1", Firstname: "Patrick", Lastname: "Schadler"})
	people = append(people, Person{ID: "2", Firstname: "Frodo", Lastname: "Beutlin"})
	people = append(people, Person{ID: "3", Firstname: "Samweis", Lastname: "Gamdschie"})
}

func main() {
	provideTestPeople()

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/addTwoNumber", AddTwoNumbersHandler)
	router.HandleFunc("/squareNumber", SquareHandler)
	router.HandleFunc("/divideNumber", DivideNumbersHandler)
	router.HandleFunc("/subtract", SubtractNumbersHandler)
	router.HandleFunc("/modulo", ModuloNumbersHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}
