package database

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var instance *IMongoInstance
var once sync.Once

func GetMongoInstance() *IMongoInstance {
	once.Do(func() {
		//DATABASE CONNECTION
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://Duong:09072003Tpl@quotes.kktlaqu.mongodb.net/?retryWrites=true&w=majority").
			SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		instance = &IMongoInstance{
			Client: client,
			Db:     client.Database("Quotes"),
		}

	})

	return instance
}
