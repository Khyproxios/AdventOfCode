package main

type ClearCommand struct{}

func (_ ClearCommand) Setup() {}

func (command ClearCommand) Execute(languages []Language) {}
