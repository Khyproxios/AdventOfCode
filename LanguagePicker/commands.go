package main

type Command interface {
	Setup()
	Execute(languages []Language)
}
