package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo(data string) *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	/* if err = client.Connect(ctx); err != nil {
		log.Fatal("Not Connected")
	} */

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("MongoDB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	DB = client.Database(data)
	return client

}

func Get() *mongo.Database {
	return DB
}
