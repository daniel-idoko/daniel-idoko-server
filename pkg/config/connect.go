package config

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(port string, r *mux.Router) {
	var clientOptions = os.Getenv("DB_CONNECTION_STRING")

	//** CONNECTING TO MONGODB & SERVE
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(clientOptions))
	if err != nil {
		panic(err)
	} else {
		Client = mongoClient
		fmt.Printf("Starting server ...")
		http.ListenAndServe(":"+port, r)
	}
}

func GetBlogCollection() *mongo.Collection {
	theCollection := Client.Database("goLang").Collection("bahd-since-O2-blogs")
	return theCollection
}

func GetProjectCollection() *mongo.Collection {
	theCollection := Client.Database("goLang").Collection("bahd-since-O2-projects")
	return theCollection
}

func GetNewsLetterCollection() *mongo.Collection {
	theCollection := Client.Database("goLang").Collection("bahd-since-O2-newsletter-emails")
	return theCollection
}

func GetCommentCollection() *mongo.Collection {
	theCollection := Client.Database("goLang").Collection("bahd-since-O2-comments")
	return theCollection
}
