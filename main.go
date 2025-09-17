package main

import (
	//"context"
	//DBService "profit-tracker-go/DataServices"

	DataStructures "profit-tracker-go/DataStructures"
	Interface "profit-tracker-go/Interfaces"
	"strings"
)

func main() {
	//mongoClient := DBService.SetMongoClient()

	//defer mongoClient.Client.Disconnect(context.Background())
	prompt := Interface.Prompt{
		PathStack: DataStructures.CreateStack(),
	}

	result := "0. Main Menu"
	for {
		if prompt.PathStack.Len() == 0 {
			prompt.PathStack.Push(result)
		}

		result = Interface.RunPrompt(result, &prompt)

		if strings.Split(result, ". ")[1] != "Prev" {
			prompt.PathStack.Push(result)
		}

		if result == "4. Exit" {
			break
		}
	}
}
