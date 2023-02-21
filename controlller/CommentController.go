package controllers

import (
	"encoding/json"
	"github.com/Dwata-Tech/golang-test-task/model"
	"github.com/Dwata-Tech/golang-test-task/service"
	"github.com/Dwata-Tech/golang-test-task/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetCommentsList(w http.ResponseWriter, r *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(r)["articleId"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	res, statusCode, err := service.GetCommentListService(articleId)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, statusCode, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ar model.Comment

	err := json.NewDecoder(r.Body).Decode(&ar)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, statusCode, err := service.CreateCommentService(ar)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, statusCode, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		logrus.Error("Error: " + err.Error())
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
