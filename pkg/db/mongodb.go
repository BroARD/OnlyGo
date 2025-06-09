package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitMongoDB() (*mongo.Client, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	cred := options.Credential{
		Username: "mongo",
		Password: "yourpassword",
		AuthSource: "admin",
	}

	clientOpts := options.Client().ApplyURI("mongodb://mongo:27017").SetAuth(cred)

	client, err := mongo.Connect(clientOpts)

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
    }

	log.Println("Подключение к MongoDB успешно установлено")

	return client, err
}
