package controllers

import (
	"encoding/json"
	"github.com/Dwata-Tech/golang-test-task/database"
	"github.com/Dwata-Tech/golang-test-task/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ar model.Article
	json.NewDecoder(r.Body).Decode(&ar)
	database.Instance.Create(&ar)
	json.NewEncoder(w).Encode(ar)
}

func GetArticleDetails(w http.ResponseWriter, r *http.Request) {
	articleId := mux.Vars(r)["id"]
	if checkIfProductExists(articleId) == false {
		json.NewEncoder(w).Encode("Article Not Found!")
		return
	}
	var article model.Article
	database.Instance.First(&article, articleId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func GetArticlesList(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	// Set default values if not provided
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}

	// Calculate offset and limit
	offset := (page - 1) * pageSize
	limit := pageSize

	var articles []model.Article
	database.Instance.Offset(offset).Limit(limit).Find(&articles)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}

func checkIfProductExists(id string) bool {
	var article model.Article
	database.Instance.First(&article, id)
	if article.ID == 0 {
		return false
	}
	return true
}
