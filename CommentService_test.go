package main

import (
	"github.com/Dwata-Tech/golang-test-task/config"
	"github.com/Dwata-Tech/golang-test-task/database"
	"github.com/Dwata-Tech/golang-test-task/model"
	"github.com/Dwata-Tech/golang-test-task/service"
	"net/http"
	"testing"
)

/**
GET Comments API
Success Example
*/
func TestGetCommentServiceSuccess(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.GetCommentListService(1)
	if i != http.StatusOK {
		t.Errorf("got %q, wanted %q", i, http.StatusOK)
	}
}

/**
GET Comments API
Not Found Example
*/
func TestGetCommentServiceArticleNotFound(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.GetCommentListService(100)
	if i != http.StatusBadRequest {
		t.Errorf("got %q, wanted %q", i, http.StatusBadRequest)
	}
}

/**
Create Comment API
Success Example
*/
func TestCreateCommentServiceArticleSuccess(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.CreateCommentService(model.Comment{Nickname: "Nickname", Content: "This is content", ArticleID: 1})
	if i != http.StatusOK {
		t.Errorf("got %q, wanted %q", i, http.StatusOK)
	}
}

/**
Create Comment API
Success Example
*/
func TestCreateCommentServiceBadRequest1(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.CreateCommentService(model.Comment{Nickname: "Nickname", ArticleID: 1})
	if i != http.StatusBadRequest {
		t.Errorf("got %q, wanted %q", i, http.StatusBadRequest)
	}
}

/**
Create Comment API
Success Example
*/
func TestCreateCommentServiceBadRequest2(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.CreateCommentService(model.Comment{Content: "This is content", ArticleID: 1})
	if i != http.StatusBadRequest {
		t.Errorf("got %q, wanted %q", i, http.StatusBadRequest)
	}
}

/**
Create Comment API
Bad request Example
*/
func TestCreateCommentServiceArticleNotFound(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.CreateCommentService(model.Comment{Nickname: "Nickname", Content: "This is content", ArticleID: 100})
	if i != http.StatusBadRequest {
		t.Errorf("got %q, wanted %q", i, http.StatusBadRequest)
	}
}
