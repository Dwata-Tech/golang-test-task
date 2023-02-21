package main

import (
	"fmt"
	"github.com/Dwata-Tech/golang-test-task/config"
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
	config.LoadAppConfig()

	// Initialize Database
	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	//init the rabbitMQ
	//rabbitmq.Connect()

	////init consumer
	//rabbitmq.StartConsumer()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/articles", controllers.GetArticlesList).Methods("GET")
	router.HandleFunc("/api/articles/{id}", controllers.GetArticleDetails).Methods("GET")
	router.HandleFunc("/api/article", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/api/comments/{articleId}", controllers.GetCommentsList).Methods("GET")
	router.HandleFunc("/api/comments", controllers.CreateComment).Methods("POST")
}
