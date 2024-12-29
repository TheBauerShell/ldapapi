package services

import (
	"encoding/json"
	"errors"
	"go-webapi/internal/models"
	"os"
	"strconv"
	"sync"
)

var persons = make(map[string]models.Person)
var mu sync.Mutex
var dataFile = "data/persons.json"

func init() {
	LoadData()
}

func LoadData() error {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(&persons)
}
func saveData() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(persons)
}

func GetPersons() []models.Person {
	mu.Lock()
	defer mu.Unlock()
	list := make([]models.Person, 0, len(persons))
	for _, person := range persons {
		list = append(list, person)

	}
	return list
}

func GetPerson(id string) (models.Person, error) {

	mu.Lock()
	defer mu.Unlock()

	if person, exists := persons[id]; exists {
		return person, nil
	}
	return models.Person{}, errors.New("person not found!")

}
func CreatePerson(person models.Person) {
	mu.Lock()
	defer mu.Unlock()
	// Generate a new ID
	newID := strconv.Itoa(len(persons) + 1)
	person.ID = newID
	persons[person.ID] = person
	saveData()
}

func UpdatePerson(id string, updatedPerson models.Person) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := persons[id]; exists {
		persons[id] = updatedPerson
		saveData()
		return nil
	}
	return errors.New("person not found")
}

func DeletePerson(id string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := persons[id]; exists {
		delete(persons, id)
		saveData()
		return nil
	}
	return errors.New("person not found")
}

func ClearPersons() {
	mu.Lock()
	defer mu.Unlock()
	persons = make(map[string]models.Person)
	saveData()
}
