package main

import "fmt"

type ListCommand struct{}

func (_ ListCommand) Setup() {}

func (command ListCommand) Execute(languages []Language) {
	if len(languages) == 0 {
		fmt.Println("No languages present")
	}
}
