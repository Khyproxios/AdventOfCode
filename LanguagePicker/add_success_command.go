package main

type AddSuccessCommand struct{ Id uint8 }

func (_ AddSuccessCommand) Setup() {}

func (command AddSuccessCommand) Execute(languages []Language) {}
