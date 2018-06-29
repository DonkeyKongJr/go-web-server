package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

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

func DivideNumbersHandler(w http.ResponseWriter, r *http.Request) {
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

		fmt.Fprint(w, DivideNumbers(calcNumbers.Number1, calcNumbers.Number2))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func SubtractNumbersHandler(w http.ResponseWriter, r *http.Request) {
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

		fmt.Fprint(w, SubtractTwoNumbers(calcNumbers.Number1, calcNumbers.Number2))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func ModuloNumbersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		var calcNumbers CalcWholeNumbers
		if err := json.Unmarshal([]byte(body), &calcNumbers); err != nil {
			panic(err)
		}

		fmt.Fprint(w, ModuloNumber(calcNumbers.Number1, calcNumbers.Number2))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func MultiplyNumbersHandler(w http.ResponseWriter, r *http.Request) {
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

		fmt.Fprint(w, MultiplyTwoNumbers(calcNumbers.Number1, calcNumbers.Number2))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

