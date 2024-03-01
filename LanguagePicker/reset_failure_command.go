package main

type ResetFailureCommand struct{ Id uint8 }

// type ResetFailureCommand struct{}
func (_ ResetFailureCommand) Setup() {}

func (command ResetFailureCommand) Execute(languages []Language) {}
