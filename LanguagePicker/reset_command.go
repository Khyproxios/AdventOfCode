package main

type ResetCommand struct{}

func (_ ResetCommand) Setup() {}

func (command ResetCommand) Execute(languages []Language) {
}
