package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"io/ioutil"
	"fmt"
)

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func AddTwoNumbers(firstNumber float32, secondNumber float32) float32{
	return firstNumber + secondNumber;
}

func SquareNumber(number float32) float32{
	return number * number;
}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
}

type CalcNumbers struct {
	Number1 float32
	Number2 float32
}

// PostHandler converts post request body to string
func AddTwoNumbersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		var calcNumbers CalcNumbers
		if err := json.Unmarshal([]byte(body), &calcNumbers); err != nil {
			panic(err)
		}

		fmt.Fprint(w, AddTwoNumbers(calcNumbers.Number1, calcNumbers.Number2))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// PostHandler converts post request body to string
func SquareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		var calcNumbers CalcNumbers
		if err := json.Unmarshal([]byte(body), &calcNumbers); err != nil {
			panic(err)
		}

		fmt.Fprint(w, SquareNumber(calcNumbers.Number1))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
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

	log.Fatal(http.ListenAndServe(":8000", router))
}
