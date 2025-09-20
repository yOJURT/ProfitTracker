package DataServices

import (
	"encoding/json"
	"os"
)

type JsonReader struct {
	FilePath string
	//TODO: Add json reader here
}

type DataBaseSettings struct {
	ConnectionString string
	DatabaseName     string
	CollectionName   string
}

func GetDatabaseSettings() *DataBaseSettings {
	file, err := os.Open("C:\\Users\\naviv\\source\\repos\\ProfitTracker\\DataServices\\settings.json")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var dbSettings DataBaseSettings
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dbSettings)
	if err != nil {
		panic(err)
	}

	return &dbSettings
}

func SetDataBaseSettings(dbsettings *DataBaseSettings) {
	newSettings, err := json.MarshalIndent(dbsettings, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("C:\\Users\\naviv\\source\\repos\\ProfitTracker\\DataServices\\settings.json", newSettings, 0644)
	if err != nil {
		panic(err)
	}

}
