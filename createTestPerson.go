package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func main() {
	people := make(map[string]Person)

	for i := 1; i <= 4000; i++ {
		id := strconv.Itoa(i)
		person := Person{
			ID:        id,
			FirstName: fmt.Sprintf("FirstName%d", i),
			LastName:  fmt.Sprintf("LastName%d", i),
			Email:     fmt.Sprintf("user%d@example.com", i),
			Phone:     fmt.Sprintf("+100000000%04d", i),
		}
		people[id] = person
	}

	file, err := os.Create("persons.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(people)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
