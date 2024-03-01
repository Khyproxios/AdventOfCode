package main

import (
	"fmt"
	"os"
	"path"
)

type Language struct {
	Id        uint8
	Name      string
	Successes uint8
	Failures  uint8
}

func main() {
	directory, error := os.UserHomeDir()

	if error != nil {
		fmt.Println(error)

		return
	}

	directory = path.Join(directory, ".khyproxios", "lang-picker")
	error = os.MkdirAll(directory, 0777)

	if error != nil {
		fmt.Println(error)

		return
	}
	directory = path.Join(directory, "languages.json")

	commands := map[string]Command{
		"--help":            HelpCommand{},
		"--add-success":     AddSuccessCommand{},
		"--add-failure":     AddFailureCommand{},
		"--list":            ListCommand{},
		"--add":             AddLanguageCommand{},
		"--clear":           ClearCommand{},
		"--reset-successes": ResetSuccessCommand{},
		"--reset-failures":  ResetFailureCommand{},
	}

	if len(os.Args) == 1 {
		fmt.Println("No command entered; printing help below.")

		commands["--help"].Execute([]Language{})

		return
	}

	commandArgument := os.Args[1]

	command := commands[commandArgument]

	if command == nil {
		fmt.Println("Unknown command entered. Pass '--help' to the program to see command options.")
	}

	command.Setup()

	verifyStore(directory)

	data, error := loadData(directory)

	if error != nil {
		fmt.Println("Error loading store:", error)
	}

	languages, error := parseLanguages(data)

	command.Execute(languages)
}

func oldMain() {
	filename := os.Args[1]

	data, error := loadData(filename)

	if error != nil {
		fmt.Println(error)
	}

	languages, error := parseLanguages(data)

	fmt.Println("Language parsed:", languages)
}
