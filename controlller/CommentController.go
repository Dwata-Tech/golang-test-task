package controllers

import (
	"encoding/json"
	"github.com/Dwata-Tech/golang-test-task/database"
	"github.com/Dwata-Tech/golang-test-task/model"
	"github.com/gorilla/mux"
	"net/http"
)

func GetCommentsList(w http.ResponseWriter, r *http.Request) {
	articleId := mux.Vars(r)["articleId"]
	if checkIfProductExists(articleId) == false {
		json.NewEncoder(w).Encode("Article Not Found!")
		return
	}
	var comments []model.Comment
	database.Instance.Find(&comments)
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
