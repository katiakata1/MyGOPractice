package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name  string
	Email string
}

func main() {
	const templateString = "Hello, {{.Name}}! Your email is {{.Email}}"

	user := User{
		Name:  "Frodo",
		Email: "jan@gmail.com",
	}

	// Parse the HTML template string
	tmpl, err := template.New("test").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	// Render the template with the data (user) and output it to a file
	file, err := os.Create("user_output.md")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = tmpl.Execute(file, user)
	if err != nil {
		log.Fatal(err)
	}

}
