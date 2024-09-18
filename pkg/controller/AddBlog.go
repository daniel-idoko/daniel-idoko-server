package controller

import (
	"bahd-since-O2/pkg/config"
	"bahd-since-O2/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type BlogPost struct {
	Number       int64    `json:"number"`
	Title        string   `json:"title"`
	Date         string   `json:"date"`
	Img          string   `json:"img"`
	Imgalt       string   `json:"imgalt"`
	Imgsource    string   `json:"imgsource"`
	Category     string   `json:"category"`
	Readduration string   `json:"readduration"`
	Tags         []string `json:"tags"`
	Views        int64    `json:"views"`
	Smallbody    string   `json:"Smallbody"`
	Body         string   `json:"body"`
}

func AddBlogHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	var newBlog BlogPost
	json.NewDecoder(r.Body).Decode(&newBlog)

	if newBlog.Body != "" {
		GetBlogCollection := config.GetBlogCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := GetBlogCollection.InsertOne(ctx, newBlog)
		if err != nil {
			panic(err)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	}
}
