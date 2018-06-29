package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)


func main() {

	router := mux.NewRouter()
	router.HandleFunc("/addTwoNumber", AddTwoNumbersHandler)
	router.HandleFunc("/squareNumber", SquareHandler)
	router.HandleFunc("/divideNumber", DivideNumbersHandler)
	router.HandleFunc("/subtract", SubtractNumbersHandler)
	router.HandleFunc("/modulo", ModuloNumbersHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}
