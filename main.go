package main

import (
	"fmt"
	controllers "github.com/Dwata-Tech/golang-test-task/controlller"
	"github.com/Dwata-Tech/golang-test-task/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/articles", controllers.GetArticlesList).Methods("GET")
	router.HandleFunc("/api/articles/{id}", controllers.GetArticleDetails).Methods("GET")
	router.HandleFunc("/api/article", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/api/comments/{articleId}", controllers.GetCommentsList).Methods("GET")
	router.HandleFunc("/api/comments", controllers.CreateComment).Methods("POST")
}
