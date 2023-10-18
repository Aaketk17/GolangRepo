package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Aaketk17/MongoAPIs/model"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://athavantheivendram:kAkA9GSEzkgyCbQl@cluster0.vq0psh5.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// Connect with mongoDB
// ! Need to check PRIMITIVES and bson.M vs bson.D
func init() {

	//? Method 01

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// * Here context will valid until 10sec. context It allows you to pass cancellation signals, deadlines,
	// * and other request-scoped values across API boundaries and between goroutines.

	// ! It's important to note that the context timeout doesn't close the established connection after 10 seconds; it simply limits how long the connection attempt can take.
	// ! Once the connection is successfully established, it will persist until you explicitly close it, or until the MongoDB server disconnects.

	//? Method 02

	// Client option
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		fmt.Println("Error is:", err)
		log.Fatal(err)
	}

	// * Here context is set to TODO, we will use TODO if we are not clear about which context to use
	// * If we use context.Background() It is never canceled, has no values, and has no deadline.
	// * It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
	// * It will keeps on happening in the background

	fmt.Println("MongoDB Connection established")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready for use")
}

//? MongoDB helpers -- separate file

// inset one record
func insertOneMovie(movie model.NetFlix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update one record
func updateOneMovie(movieId string) {
	//? converting the id to objectid which mongo can understand
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)

}

// delete one record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got delete with delete countL ", deleteCount.DeletedCount)
}

// delete all record
func deleteAllMovies() int64 {
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All records were deleted", deleteCount.DeletedCount)
	return deleteCount.DeletedCount

}

// get all movies
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies

}

//? Controllers -- separate file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.NetFlix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	responseString := fmt.Sprintf("Movie with id %s is marked as watched", params["id"])
	json.NewEncoder(w).Encode(responseString)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	responseString := fmt.Sprintf("Movie with id %s is deleted", params["id"])
	json.NewEncoder(w).Encode(responseString)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode(count)
}
