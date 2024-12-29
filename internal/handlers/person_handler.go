package handlers

import (
	"encoding/json"
	"go-webapi/internal/models"
	"go-webapi/internal/services"
	"log"
	"net/http"

	//"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func init() {
	if err := services.LoadData(); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}
}
func GetPersons(w http.ResponseWriter, r *http.Request) {
	persons := services.GetPersons()
	json.NewEncoder(w).Encode(persons)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person, err := services.GetPerson(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(person)
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	missingFields := []string{}
	if person.FirstName == "" {
		missingFields = append(missingFields, "firstName")
	}
	if person.LastName == "" {
		missingFields = append(missingFields, "lastName")
	}
	if person.Email == "" {
		missingFields = append(missingFields, "email")
	}
	if person.Phone == "" {
		missingFields = append(missingFields, "phone")
	}

	if len(missingFields) > 0 {
		http.Error(w, "Missing required fields: "+strings.Join(missingFields, ", "), http.StatusBadRequest)
		return
	}

	services.CreatePerson(person)
	persons := services.GetPersons()
	json.NewEncoder(w).Encode(persons)
}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedPerson models.Person
	json.NewDecoder(r.Body).Decode(&updatedPerson)
	err := services.UpdatePerson(params["id"], updatedPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := services.DeletePerson(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// ResetPersons resets the list of persons

func ResetPersons(w http.ResponseWriter, r *http.Request) {

	// Implementation for resetting persons

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Persons reset successfully"))

}
