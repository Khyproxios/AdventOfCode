package main

import "fmt"

type HelpCommand struct{}

func (_ HelpCommand) Setup() {}

func (command HelpCommand) Execute(_ []Language) {
	fmt.Println(`
		lang-picker is a tool for picking a programming language, tracking their usages
			and the amount of successes and failures.

		Usage:
			lang-picker <command> [arguments]

		Commands available:

			--help					print out this message about usage
			--add-success 			add to success count for a language
			--add-failure 			add to failure count for a language
			--list					print out the current languages with their metrics
			--add					add one or more languages
			--clear					clear the languages
			--reset-successes		reset the successes
			--reset-failures		reset the failures
		`)
}
