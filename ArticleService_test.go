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
CREATE Article API
should return bad request if Nickname Missing
*/
func TestCreateArticleServiceBadRequest1(t *testing.T) {

	_, i, _ := service.CreateArticleService(model.Article{Title: "Title", Content: "This is content"})
	if i != http.StatusBadRequest {
		t.Errorf("got %q, wanted %q", i, http.StatusBadRequest)
	}
}

/**
CREATE Article API
should return bad request if Title Missing
*/
func TestCreateArticleServiceBadRequest2(t *testing.T) {

	_, i, _ := service.CreateArticleService(model.Article{Nickname: "Vaibhav", Content: "This is content"})
	if i != http.StatusBadRequest {
		t.Errorf("got %q, wanted %q", i, http.StatusBadRequest)
	}
}

/**
CREATE Article API
should return bad request if Content Missing
*/
func TestCreateArticleServiceBadRequest3(t *testing.T) {

	_, i, _ := service.CreateArticleService(model.Article{Nickname: "Vaibhav", Title: "This is title"})
	if i != http.StatusBadRequest {
		t.Errorf("got %q, wanted %q", i, http.StatusBadRequest)
	}
}

/**
CREATE Article API
Success Example
*/
func TestCreateArticleServiceSuccess(t *testing.T) {
	_, i, _ := service.CreateArticleService(model.Article{Nickname: "Vaibhav", Title: "This is title", Content: "This is content"})
	if i != http.StatusOK {
		t.Errorf("got %q, wanted %q", i, http.StatusOK)
	}
}

/**
GET Article API
Success Example
*/
func TestGetArticleServiceSuccess(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.GetArticleService(1)
	if i != http.StatusOK {
		t.Errorf("got %q, wanted %q", i, http.StatusOK)
	}
}

/**
GET Article API
Not Found Example
*/
func TestGetArticleServiceNotFound(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.GetArticleService(100)
	if i != http.StatusBadRequest {
		t.Errorf("got %q, wanted %q", i, http.StatusBadRequest)
	}
}

/**
GET Article List API
Success Example
*/
func TestGetArticleListServiceSuccess(t *testing.T) {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)

	_, i, _ := service.GetArticleListService(1, 10)
	if i != http.StatusOK {
		t.Errorf("got %q, wanted %q", i, http.StatusOK)
	}
}
