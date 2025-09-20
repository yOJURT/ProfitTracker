package worker

import (
	DataService "profit-tracker-go/DataServices"
	DataStructures "profit-tracker-go/DataStructures"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PointersToUse struct {
	DatabaseSettings *DataService.DataBaseSettings
	MongoClient      *mongo.Client
	MongoCollection  *mongo.Collection
	PathStack        *DataStructures.Stack
	MainMenuItems    []string
}

func Run() *PointersToUse {
	var pointers PointersToUse
	DataService.SetMongoClient("mongodb://localhost:27017")
	pointers.MainMenuItems = []string{"1. View Data", "2. Insert Data", "3. Settings", "4. Connect To Mongo", "5. Exit"}
	pointers.DatabaseSettings = DataService.GetDatabaseSettings()
	pointers.PathStack = DataStructures.CreateStack()
	pointers.MongoClient = nil
	pointers.MongoCollection = nil
	return &pointers
}
