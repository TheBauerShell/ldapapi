package main

import (
	"go-webapi/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/persons", handlers.GetPersons).Methods("GET")
	r.HandleFunc("/persons/{id}", handlers.GetPerson).Methods("GET")
	r.HandleFunc("/person", handlers.CreatePerson).Methods("POST")
	r.HandleFunc("/person/{id}", handlers.UpdatePerson).Methods("PUT")
	r.HandleFunc("/person/{id}", handlers.DeletePerson).Methods("DELETE")
	r.HandleFunc("/reset-persons", handlers.ResetPersons).Methods("POST")

	log.Println("Starting Server on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
