package controller

import (
	"bahd-since-O2/pkg/config"
	"bahd-since-O2/pkg/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NewsLetterResponse struct {
	Success     bool `json:"success"`
	MessageCode int  `json:"message_code"`
}

func AddNewsLetterSubcriber(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	GetNewsLetterCollection := config.GetNewsLetterCollection()

	filter := bson.M{"email": email}
	// Data to update or insert
	update := bson.M{
		"$set": bson.M{
			"email":      email,
			"isverified": false,
		},
	}
	// Upsert option - if the document is not found, it will be inserted
	opts := options.Update().SetUpsert(true)
	// Perform the upsert operation
	result, err := GetNewsLetterCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal("Error performing upsert:", err)
	}

	var newNewsLetterResponse NewsLetterResponse
	if result.MatchedCount > 0 {
		newNewsLetterResponse.Success = true
		newNewsLetterResponse.MessageCode = 2
		json.NewEncoder(w).Encode(newNewsLetterResponse)
	} else if result.UpsertedCount > 0 {
		newNewsLetterResponse.Success = true
		newNewsLetterResponse.MessageCode = 1
		json.NewEncoder(w).Encode(newNewsLetterResponse)
	} else {
		newNewsLetterResponse.Success = false
		newNewsLetterResponse.MessageCode = 3
		json.NewEncoder(w).Encode(newNewsLetterResponse)
	}
}
