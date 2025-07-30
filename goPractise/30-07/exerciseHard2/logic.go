// 1. Normalise names Title-cased: "alice Johnson" → "Alice Johnson"
// 2. Standardize the Email "Alice.Johnson@company.COM" → "alice.johnson@company.com"
// 3. Generate a Unique Username "Alice Johnson" → "ajohnson"
// 4. Group by Environment. Create a map[string][]string where: Each key is a normalized environment ("dev", "staging", "production", etc.), Each value is a list of Usernames with access to that environment
// 5. Output Print the grouped access map as indented JSON.

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
	Email       string
	Username    string
}

func process(employees []Employee) ([]Employee, map[string][]string) {

	newEmployee := []Employee{}
	groupByEnv := map[string][]string{}

	// Normalising names - removing extra space, capital letter
	for _, employee := range employees {
		splittedNames := strings.Fields(employee.Name)
		for i, name := range splittedNames {
			correctCaseNames := strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
			splittedNames[i] = correctCaseNames
		}
		correctNames := strings.Join(splittedNames, " ")
		employee.Name = correctNames

		// Standardilise emails - Trim whitespace, Convert to lowercase
		employee.Email = strings.TrimSpace(strings.ToLower(employee.Email))

		// Generate unique username
		if len(splittedNames) >= 2 {
			firstName := splittedNames[0]
			userName := string(firstName[0]) + splittedNames[len(splittedNames)-1]
			employee.Username = strings.ToLower(userName)
		}

		// groupByEnv := map[string][]string{}
		arrayEnvs := strings.Split(employee.AccessLevel, ",")
		for _, env := range arrayEnvs {
			correctEnv := strings.ToLower(strings.TrimSpace(env))
			if _, exist := groupByEnv[correctEnv]; exist {
				groupByEnv[correctEnv] = append(groupByEnv[correctEnv], employee.Username)
			} else {
				groupByEnv[correctEnv] = []string{employee.Username}
			}
		}

		newEmployee = append(newEmployee, employee)
	}

	return newEmployee, groupByEnv

}

func marshalJson(employees map[string][]string) {
	data, err := json.MarshalIndent(employees, "", " ")
	if err != nil {
		fmt.Println("error marhsaling from map to bytes: ", err)
		return
	}

	if err := os.WriteFile("output.json", data, 0644); err != nil {
		fmt.Println("error writing to file: ", err)
		return
	}
}

func main() {
	employees := []Employee{
		{Name: "alice   Johnson", AccessLevel: "dev, staging ,production", Email: "Alice.Johnson@company.COM"},
		{Name: "Bob smith", AccessLevel: "DEV,Production", Email: "bob.SMITH@company.com "},
		{Name: "Clara Olivia West", AccessLevel: "staging, qa", Email: " clara.west@COMPANY.com"},
		{Name: "david miles", AccessLevel: "dev", Email: "david.miles@company.com"},
	}

	emp, newmap := process(employees)
	fmt.Println(emp, newmap)
	marshalJson(newmap)
}
