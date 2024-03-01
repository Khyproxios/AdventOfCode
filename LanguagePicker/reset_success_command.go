package main

type ResetSuccessCommand struct{ Id uint8 }

// type ResetFailureCommand struct{}
func (_ ResetSuccessCommand) Setup() {}

func (command ResetSuccessCommand) Execute(languages []Language) {}
