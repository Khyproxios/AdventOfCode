package main

type AddFailureCommand struct{ Id uint8 }

func (_ AddFailureCommand) Setup() {}

func (command AddFailureCommand) Execute(languages []Language) {}
