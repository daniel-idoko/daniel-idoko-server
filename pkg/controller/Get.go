package controller

import (
	"bahd-since-O2/pkg/config"
	"bahd-since-O2/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type BlogNumber struct {
	Number int64
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.GetBlogCollection()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			fmt.Println(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func GetProjectHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	theCollection := config.GetProjectCollection()
	cursor, err := theCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	} else {
		var content []bson.M
		if err = cursor.All(ctx, &content); err != nil {
			fmt.Println(err)
		} else {
			json.NewEncoder(w).Encode(content)
		}
	}
}

// func GetSingleBlogObject(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userID := vars["blogID"]
// }

// func GetBlogBody(w http.ResponseWriter, r *http.Request) {
// 	utils.EnableCors(w, r)

// 	var newBlogNumber BlogNumber
// 	json.NewDecoder(r.Body).Decode(&newBlogNumber)

// 	fmt.Println(newBlogNumber)
// 	if newBlogNumber.Number != 0 {
// 		GetBlogCollection := config.GetBlogCollection()
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 		var content bson.M
// 		filter := bson.D{{Key: "number", Value: newBlogNumber.Number}}
// 		result := GetBlogCollection.FindOne(ctx, filter)
// 		err := result.Decode(&content)
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			json.NewEncoder(w).Encode(content)
// 		}
// 	}
// }
