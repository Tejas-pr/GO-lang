package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mangodbb/model"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName         = "netflix"
	collectionName = "watchlist"
	collections    *mongo.Collection
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	connectionString := os.Getenv("DB_URL")
	if connectionString == "" {
		log.Fatal("DB_URL is not set in environment variables")
	}
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected!!")
	collections = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance ready!!")
}

// Insert one movie
func insertOneMovie(movie model.Netflix) (interface{}, error) {
	insert, err := collections.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserted one movie with id: ", insert.InsertedID)
	return insert.InsertedID, nil
}

// Update one movie
func updateOneMovie(movieId string) (int64, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collections.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	fmt.Println("Modified count: ", res.ModifiedCount)
	return res.ModifiedCount, nil
}

// Delete one movie
func deleteOneMovie(movieID string) (int64, error) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": id}

	res, err := collections.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	fmt.Println("Deleted count: ", res.DeletedCount)
	return res.DeletedCount, nil
}

// Delete all movies
func deleteAllMovie() int64 {
	res, err := collections.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ", res.DeletedCount)
	return res.DeletedCount
}

// read
func getAllMovies() []primitive.M {
	cur, err := collections.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// acutall controllers - file
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
