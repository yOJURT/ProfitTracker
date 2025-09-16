package main

import (
	"context"
	DBService "profit-tracker-go/DataServices"
	Interface "profit-tracker-go/Interfaces"
)

func main() {
	mongoClient := DBService.SetMongoClient()

	defer mongoClient.Client.Disconnect(context.Background())

	for {
		result := Interface.GetUserInput()
		if result == "4. Exit" {
			break
		} else {
			//TODO: Call and write parser to handle user selection
		}
	}
}
