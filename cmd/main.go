package main

import (
	"bahd-since-O2/pkg/config"
	"bahd-since-O2/pkg/controller"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Only for local
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// 	panic(err)
	// }

	r := mux.NewRouter()

	r.HandleFunc("/post-blog", controller.AddBlogHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/get-blogs", controller.GetAllHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/post-project", controller.AddProjectHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/get-projects", controller.GetProjectHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/news-letter-post", controller.AddNewsLetterSubcriber).Methods("POST", "OPTIONS")
	r.HandleFunc("/send-comment", controller.PostComments).Methods("POST", "OPTIONS")
	// r.HandleFunc("/blog/{blogID}", controller.GetSingleBlogObject).Methods("GET", "OPTIONS")

	port := os.Getenv("PORT")
	config.Connect(port, r)
}
