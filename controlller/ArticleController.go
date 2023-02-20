package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Dwata-Tech/golang-test-task/database"
	"github.com/Dwata-Tech/golang-test-task/model"
	"gorm.io/gorm"
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
	articleId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	if checkIfProductExists(articleId) == false {
		json.NewEncoder(w).Encode("Article Not Found!")
		return
	}
	var article model.Article
	result := database.Instance.First(&article, articleId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	} else if result.Error != nil {
		fmt.Println(result.Error.Error())
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
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

func checkIfProductExists(id int) bool {
	var article model.Article
	database.Instance.First(&article, id)
	if article.ID == 0 {
		return false
	}
	return true
}
