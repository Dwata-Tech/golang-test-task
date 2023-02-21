package controllers

import (
	"encoding/json"
	"github.com/Dwata-Tech/golang-test-task/model"
	"github.com/Dwata-Tech/golang-test-task/service"
	"github.com/Dwata-Tech/golang-test-task/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//parse request and handle error
	var ar model.Article
	err := json.NewDecoder(r.Body).Decode(&ar)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//handle request into service layer
	res, statusCode, err := service.CreateArticleService(ar)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, statusCode, err.Error())
		return
	}

	//parse to response json and handle error
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func GetArticleDetails(w http.ResponseWriter, r *http.Request) {

	articleId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	res, statusCode, err := service.GetArticleService(articleId)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, statusCode, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
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

	res, statusCode, err := service.GetArticleListService(page, pageSize)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, statusCode, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
