package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Defined names templates for different types of
	templates := map[string]string{
		"welcome":      "Hello, {{.name}}",
		"notification": "{{.name}}, you have a new notification: {{.notification}}",
		"error":        "Oops, an error occurred: {{.errorMessage}}",
	}

	// Parse the templates
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu: ")
		fmt.Println("1. Join")
		fmt.Println("2. Get notifications")
		fmt.Println("3. Get error log")
		fmt.Println("4. Exit")
		fmt.Println("Enter your choice:")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]any{
				"name": name,
			}
		case "2":
			fmt.Println("Enter the notification:")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]any{
				"name": name,
				"notification": notification,
			}
		case "3":
			tmpl = parsedTemplates["error"]
			data = map[string]any{
				"name": name,
				"errorMessage": "An error occurred",
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
			continue
		}
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println(err)
		}
	}

}
