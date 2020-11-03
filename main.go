package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//LoadEnv loads variables from .env
func LoadEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	fmt.Println("Starting server...")

	username := LoadEnv("USER")
	password := LoadEnv("PASSWORD")
	dbname := LoadEnv("DBNAME")

	fmt.Printf("Username: %s \n", username)
	fmt.Printf("Connecting to: %s \n", dbname)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// if cancel != nil {
	// 	fmt.Println("Ctx:", cancel)
	// 	// log.Fatal(cancel) //Do not error out
	// }
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://"+username+":"+password+"@cluster0.ing9e.mongodb.net/"+dbname+"?retryWrites=true&w=majority"))
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("client.Connect:", err)
		// log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Ping error")
		log.Fatal(err)
	}
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	fmt.Println("Database error")
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	//Establishing handles
	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")

	//Insert document to podcasts Collection
	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "Some MongoDB Podcast"},
		{Key: "author", Value: "This Guy"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastResult.InsertedID)

	//Insert multiple documents
	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode 1"},
			{"description", "The first ep."},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode 2"},
			{"description", "The second ep."},
			{"duration", 30},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodeResult.InsertedIDs)
}
