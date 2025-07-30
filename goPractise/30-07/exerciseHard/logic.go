// So you need to do the following:
// First, you need to add enabled field, and if a person has production in AccessLevel I need to set enabled to true, otherwise it's false
// Second, you need to remove any middle names in name
// Third, you need to create an ID which includes first letter of the name and full surname
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Employee struct {
	Name        string
	AccessLevel string
	Enabled     bool
	Id          string
}

func process(employees []Employee) []Employee {
	newEmployees := []Employee{}

	// Enabling based on if production exists
	for _, employee := range employees {
		environments := strings.Split(employee.AccessLevel, ",")
		for _, environment := range environments {
			environment = strings.ToLower(strings.TrimSpace(environment))

			if environment == "production" {
				employee.Enabled = true
			}
		}
		splittedNames := strings.Fields(strings.ToLower(employee.Name))

		if len(splittedNames) > 2 {
			shortNames := splittedNames[0] + " " + splittedNames[len(splittedNames)-1]
			employee.Name = shortNames
		}

		if len(splittedNames) >= 2 {
			firstName := splittedNames[0]
			firstLetter := firstName[0]
			secondName := splittedNames[len(splittedNames)-1]
			id := string(firstLetter) + secondName
			employee.Id = id
		} else {
			employee.Id = "no id"
		}

		newEmployees = append(newEmployees, employee)

	}

	return newEmployees
}

func marshalToJson(employee []Employee) {
	data, err := json.MarshalIndent(employee, "", " ")
	if err != nil {
		fmt.Println("error marshaling to json:", err)
		return
	}

	if err := os.WriteFile("output.json", data, 0644); err != nil {
		fmt.Println("error writting to file: ", err)
		return
	}
}

func main() {
	employee := []Employee{
		{
			Name:        "John Doe",
			AccessLevel: "dev,staging,production",
		},
		{
			Name:        "Anderson Sterling",
			AccessLevel: "dev",
		},
		{
			Name:        "Janel Tiffany",
			AccessLevel: "dev,Staging",
		},
		{
			Name:        "Mary Anne Maitland",
			AccessLevel: "staging,Production",
		},
	}

	newEmployee := process(employee)
	fmt.Println(newEmployee)
	marshalToJson(newEmployee)
}
