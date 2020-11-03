package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

//LoadEnv loads variables from .env
func LoadEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// Insert documents code
// func InsertDocuments(client, ctx) {

// 	//Insert document to podcasts Collection
// 	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
// 		{Key: "title", Value: "Some MongoDB Podcast"},
// 		{Key: "author", Value: "This Guy"},
// 		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(podcastResult.InsertedID)

// 	//Insert multiple documents
// 	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
// 		bson.D{
// 			{"podcast", podcastResult.InsertedID},
// 			{"title", "Episode 1"},
// 			{"description", "The first ep."},
// 			{"duration", 25},
// 		},
// 		bson.D{
// 			{"podcast", podcastResult.InsertedID},
// 			{"title", "Episode 2"},
// 			{"description", "The second ep."},
// 			{"duration", 30},
// 		},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(episodeResult.InsertedIDs)
// }

// Use cursors to retrieve documents
// func cursors() {
// 	// ** Use cursors to find documents
// 	cursor, err := episodesCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// **Load all documents into memory
// 	// var episodes []bson.M
// 	// if err = cursor.All(ctx, &episodes); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// for _, episode := range episodes {
// 	// 	fmt.Println(episode)
// 	// }

// 	// **Load documents by batches with Next
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var episode bson.M
// 		if err = cursor.Decode(&episode); err != nil {
// 			log.Fatal(err)
// 		}
// 		// fmt.Println(episode)
// 	}

// 	// ** Find document (first?)
// 	var podcast bson.M
// 	if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Println(podcast)

// 	// ** Find document by filter (bson.M)
// 	filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration": 30})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var episodesFiltered []bson.M
// 	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Println(episodesFiltered)

// 	// ** Find document by sort and range
// 	opts := options.Find()
// 	opts.SetSort(bson.D{{"duration", 1}}) // -1 descending, 1 ascending
// 	sortCursor, err := episodesCollection.Find(ctx, bson.D{
// 		{"duration", bson.D{
// 			{"$gt", 24},
// 		}},
// 	}, opts)
// 	var episodesSorted []bson.M
// 	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(episodesSorted)
// }

// Update operations
// func Update() {
// 	// ** Update documents by id
// 	id, _ := primitive.ObjectIDFromHex("5fa0fd1b444827cf354d8973")

// 	result, err := podcastsCollection.UpdateOne(
// 		ctx,
// 		bson.M{"_id": id},
// 		bson.D{
// 			{"$set", bson.D{{"author", "That Guy"}}},
// 		},
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

// 	// ** Update documents by filter match
// 	results, err := podcastsCollection.UpdateMany(
// 		ctx,
// 		bson.M{"title": "Some MongoDB Podcast"},
// 		bson.D{
// 			{"$set", bson.D{{"author", "Dude"}}},
// 		},
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Updated %v Documents!\n", results.ModifiedCount)

// 	// ** Replace by filter
// 	resultReplace, err := podcastsCollection.ReplaceOne(
// 		ctx,
// 		bson.M{"author": "Dude"},
// 		bson.M{
// 			"title":  "The Updated Podcast",
// 			"author": "Some Dude",
// 		},
// 	)
// 	fmt.Printf("Updated %v Documents!\n", resultReplace.ModifiedCount)
// }

// Delete operations
// func Delete () {
// 	// ** Delete documents
// 	// result, err := episodesCollection.DeleteOne(ctx, bson.M{"duration": 25})
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Printf("DeleteOne removed %v documents\n", result.DeletedCount)

// 	// ** Delete ALL documents matching filter
// 	// result, err := episodesCollection.DeleteMany(ctx, bson.M{"duration": 30})
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Printf("DeleteOne removed %v documents\n", result.DeletedCount)

// 	// ** Delete entire collection
// 	// if err = podcastsCollection.Drop(ctx); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// }

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
		panic(err)
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

	var podcasts []Podcast
	podcastCursor, err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if err = podcastCursor.All(ctx, &podcasts); err != nil {
		panic(err)
	}
	fmt.Println(podcasts)

	var episodes []Episode
	episodeCursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if err = episodeCursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	fmt.Println(episodes[0].Title)

}
