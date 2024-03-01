package main

type AddLanguageCommand struct{}

func (_ AddLanguageCommand) Setup() {}

func (command AddLanguageCommand) Execute(languages []Language) {}
