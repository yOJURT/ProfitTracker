package Interfaces

import (
	"fmt"
	DataServices "profit-tracker-go/DataServices"
	DataStructures "profit-tracker-go/DataStructures"
	worker "profit-tracker-go/worker"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

type Prompt struct {
	PathStack *DataStructures.Stack
	prompt    *promptui.Select
}

func RunPrompt(prompt string, p *worker.PointersToUse) string {

	var PromptToRun Prompt
	//TODO: Make this cleaner
	if prompt != "0. Main Menu" {
		prompt = strings.Split(prompt, ". ")[1]
	}

	switch prompt {
	case "0. Main Menu":
		PromptToRun = Prompt{
			prompt: GetMainMenuInput(p.MainMenuItems),
		}
	//Main Menu Options
	case "Settings":
		PromptToRun = Prompt{
			prompt: GetSettingsInput(),
		}
	case "View Data":
		if p.MongoClient == nil {
			ClearScreen()
			fmt.Println("Please connect to mongo before attempting to view data.")
			time.Sleep(3 * time.Second)
			ClearScreen()
			return "0. Main Menu"
		}

		PromptToRun = Prompt{
			prompt: GetViewDataInput(),
		}
	case "Insert Data":
		if p.MongoClient == nil {
			ClearScreen()
			fmt.Println("Please connect to mongo before attempting to insert data.")
			time.Sleep(3 * time.Second)
			ClearScreen()
			return "0. Main Menu"
		}

		PromptToRun = Prompt{
			prompt: GetInsertDataInput(),
		}
	//Connect and disconnect from database
	case "Connect To Mongo":
		p.MainMenuItems = []string{"1. View Data", "2. Insert Data", "3. Settings", "4. Disconnect From Mongo", "5. Exit"}
		mongoClient := DataServices.SetMongoClient(p.DatabaseSettings.ConnectionString)
		p.MongoClient = mongoClient.Client

		p.MongoCollection = mongoClient.Client.Database(p.DatabaseSettings.DatabaseName).Collection(p.DatabaseSettings.CollectionName)

		ClearScreen()
		fmt.Println("Sucessfully established connection to Mongo.")
		time.Sleep(3 * time.Second)
		ClearScreen()
		return "0. Main Menu"

	case "Disconnect From Mongo":
		p.MainMenuItems = []string{"1. View Data", "2. Insert Data", "3. Settings", "4. Connect To Mongo", "5. Exit"}
		DataServices.DisconnectFromMongo(p.MongoClient)

		p.MongoClient = nil
		p.MongoCollection = nil

		ClearScreen()
		fmt.Println("Sucessfully disconnected from Mongo.")
		time.Sleep(3 * time.Second)
		ClearScreen()

		return "0. Main Menu"

	//Settings Options
	case "Set Database Connection String":
		ClearScreen()
		fmt.Scan(&p.DatabaseSettings.ConnectionString)
		DataServices.SetDataBaseSettings(p.DatabaseSettings)
		ClearScreen()
		return "0. Main Menu"

	case "Set Databse Name":
		ClearScreen()
		fmt.Scan(&p.DatabaseSettings.DatabaseName)
		DataServices.SetDataBaseSettings(p.DatabaseSettings)
		ClearScreen()
		return "0. Main Menu"

	case "Set Collection":
		ClearScreen()
		fmt.Scan(&p.DatabaseSettings.CollectionName)
		DataServices.SetDataBaseSettings(p.DatabaseSettings)
		ClearScreen()
		return "0. Main Menu"

	//View Data
	case "View All Data":
		ClearScreen()
		DataServices.GetAllDocuments(p.MongoCollection)
		ClearScreen()
		return "0. Main Menu"
	}

	_, result, err := PromptToRun.prompt.Run()

	if err != nil {
		panic(err)
	}

	if strings.Split(result, ". ")[1] == "Prev" {
		p.PathStack.Pop()
		return p.PathStack.Pop()
	}

	return result
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
