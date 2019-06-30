package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const path = "configs/config.toml"

func main() {
	c := internal.NewConfig(path)
	c.Read()
	var s = internal.NewServer(connectDb(c))
	log.Fatal(http.ListenAndServe(":8070", s.Router))
}

func connectDb(c *internal.Config) (db *mongo.Database) {
	var err error
	client, err := mongo.NewClient(options.Client().ApplyURI(c.Database.Url))
	if err != nil {
		panic(err)
	}
	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		panic(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	result := client.Database(c.Database.Name)
	log.Println("Connected to MongoDB!")

	return result
}
