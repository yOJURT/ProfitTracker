package main

import (
	//"context"
	//DBService "profit-tracker-go/DataServices"
	Interface "profit-tracker-go/Interfaces"
)

func main() {
	//mongoClient := DBService.SetMongoClient()

	//defer mongoClient.Client.Disconnect(context.Background())
	result := Interface.RunPrompt("Main Menu")
	for {
		if result == "4. Exit" {
			break
		} else {
			result = Interface.RunPrompt(result)
		}
	}
}
