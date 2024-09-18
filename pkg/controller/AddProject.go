package controller

import (
	"bahd-since-O2/pkg/config"
	"bahd-since-O2/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Blog struct {
	Number      int64  `json:"number"`
	Name        string `json:"name"`
	Year        string `json:"year"`
	Disc        string `json:"disc"`
	Tech        string `json:"tech"`
	Sourcelink  int    `json:"sourcelink"`
	Articlelink string `json:"articlelink"`
	Demolink    string `json:"demolink"`
}

func AddProjectHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newPost Blog
	json.NewDecoder(r.Body).Decode(&newPost)

	theCollection := config.GetProjectCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := theCollection.InsertOne(ctx, newPost)
	if err != nil {
		panic(err)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}
