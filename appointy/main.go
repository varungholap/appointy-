package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	ID        string `json:"id"`
	Caption   string `json:"cap"`
	Image_url string `json:"url"`
	Timestamp string `json:"timestamp"`
}
type Posts []Post

var posts Post

func allPosts(w http.ResponseWriter, r *http.Request) {
	posts := Posts{
		Post{ID: "varun", Caption: "flyhigh", Image_url: "google.com", Timestamp: "hgh"},
	}
	fmt.Println("Homepage :All Posts endpoint ")
	json.NewEncoder(w).Encode(posts)
}

type Users []User

var users User

func allUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{ID: "varun", Name: "Varun", Email: "varunrajeev@gmail.com", Password: "varungholap"},
	}
	_ = users
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {

}

const message = "Homepage Instagram Appointy"

func handleRequests() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))

	})
	http.HandleFunc("/posts", allPosts)

	http.HandleFunc("/users", allUsers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://varungholap:welcome2001@cluster0.hfhrb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	handleRequests()
}
