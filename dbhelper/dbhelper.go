package dbhelper

import (
	"context"
	"fmt"
	"log"

	"github.com/aswinayyolath/goapimongoDB/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://<username>:<password>@<dbconnectionname>"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

//connect with mongoDB

// The init function is a special function that can be defined in any Go source file.
// It is automatically called by the Go runtime before the main function is executed.
// Each package can have its own init function, and it is often used for setting up package-level variables,
// performing one-time initialization tasks, or registering things that need to happen before the program starts running.
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB

	// The context package in Go is used to manage the flow of deadlines,
	// cancellations, and other request-scoped values in your program.
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkError(err)
	fmt.Println("successfully connected to mongoDB")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready!")

}

func InsertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	checkError(err)
	fmt.Println("inserted one movie with id", inserted.InsertedID)
}

func UpdateOneMovie(movieID string) {
	filter := filterMovie(movieID)
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkError(err)
	fmt.Println("Modified count", result.ModifiedCount)
}

func filterMovie(movieID string) primitive.M {
	id, err := primitive.ObjectIDFromHex(movieID)
	checkError(err)
	filter := bson.M{"_id": id}
	return filter
}

func DeleteOneMovie(movieID string) {
	filter := filterMovie(movieID)
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	checkError(err)
	fmt.Println("Movie got deleted ", deleteCount)
}

// Delete all movie

func DeleteAllMovie() {
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	checkError(err)
	fmt.Println("Number of movies deleted", deleteCount.DeletedCount)
}

// Get all Movie

func GetAllMovie() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkError(err)
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		checkError(err)
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
