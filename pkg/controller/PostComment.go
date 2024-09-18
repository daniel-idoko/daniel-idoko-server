package controller

import (
	"bahd-since-O2/pkg/config"
	"bahd-since-O2/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type CommentPost struct {
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

func PostComments(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)

	email := r.FormValue("email")
	comment := r.FormValue("comment")

	var newComment CommentPost
	newComment.Email = email
	newComment.Comment = comment

	GetCommentCollection := config.GetCommentCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := GetCommentCollection.InsertOne(ctx, newComment)
	if err != nil {
		panic(err)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}
