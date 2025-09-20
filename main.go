package main

import (
	Interface "profit-tracker-go/Interfaces"
	worker "profit-tracker-go/worker"
	"strings"
)

func main() {
	Interface.ClearScreen()
	pointers := worker.Run()
	result := "0. Main Menu"
	for {
		if pointers.PathStack.Len() == 0 {
			pointers.PathStack.Push(result)
		}

		result = Interface.RunPrompt(result, pointers)

		if strings.Split(result, ". ")[1] != "Prev" || strings.Split(result, ". ")[1] != "Main Menu" {
			pointers.PathStack.Push(result)
		}

		if result == "5. Exit" {
			break
		}
	}
}
