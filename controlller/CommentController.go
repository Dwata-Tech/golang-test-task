package controllers

import (
	"encoding/json"
	"github.com/Dwata-Tech/golang-test-task/database"
	"github.com/Dwata-Tech/golang-test-task/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetCommentsList(w http.ResponseWriter, r *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(r)["articleId"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}
	if checkIfProductExists(articleId) == false {
		json.NewEncoder(w).Encode("Article Not Found!")
		return
	}

	var comments []model.Comment
	database.Instance.Where("article_id = ?", articleId).Find(&comments)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ar model.Comment
	json.NewDecoder(r.Body).Decode(&ar)
	database.Instance.Create(&ar)
	json.NewEncoder(w).Encode(ar)
}
