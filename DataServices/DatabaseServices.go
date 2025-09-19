package DataServices

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBConnection struct {
	Client *mongo.Client
}

func SetMongoClient(connectionString string) *DBConnection {

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context, nil)
	if err != nil {
		panic(err)
	}
	return &DBConnection{Client: client}
}

func DisconnectFromMongo(mongoClient *mongo.Client) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := mongoClient.Disconnect(context); err != nil {
		panic(err)
	}
}

func InsertDocument(document bson.M, collection *mongo.Collection) {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(context, document)

	if err != nil {
		panic(err)
	}
}

func GetAllDocuments(collection *mongo.Collection) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(context, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context)

	for cursor.Next(context) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
		fmt.Println(result)
	}

	time.Sleep(5 * time.Second)

	if err := cursor.Err(); err != nil {
		panic(err)
	}
}
