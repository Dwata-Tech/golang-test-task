package service

import (
	"errors"
	"github.com/Dwata-Tech/golang-test-task/database"
	"github.com/Dwata-Tech/golang-test-task/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func GetCommentListService(articleId int) ([]model.Comment, int, error) {
	var comments []model.Comment
	var article model.Article

	result := database.Instance.First(&article, articleId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return comments, http.StatusBadRequest, errors.New("article not found")
	} else if result.Error != nil {
		logrus.Error("Error: " + result.Error.Error())
		return comments, http.StatusInternalServerError, errors.New("internal server error")
	}

	resultComments := database.Instance.Where("article_id = ?", articleId).Find(&comments)
	if resultComments.Error != nil {
		logrus.Error("Error: " + resultComments.Error.Error())
		return comments, http.StatusInternalServerError, errors.New("internal server error")
	}
	return comments, http.StatusOK, nil
}

func CreateCommentService(comment model.Comment) (model.Comment, int, error) {

	//request validation
	if len(comment.Content) == 0 {
		return comment, http.StatusBadRequest, errors.New("content can not be null or blank")
	}

	if len(comment.Nickname) == 0 {
		return comment, http.StatusBadRequest, errors.New("nickname can not be null or blank")
	}

	if comment.ArticleID == 0 {
		return comment, http.StatusBadRequest, errors.New("article id can not be null or blank")
	}

	var article model.Article
	result := database.Instance.First(&article, comment.ArticleID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return comment, http.StatusBadRequest, errors.New("article not found")
	} else if result.Error != nil {
		logrus.Error("Error: " + result.Error.Error())
		return comment, http.StatusInternalServerError, errors.New("internal server error")
	}

	res := database.Instance.Create(&comment)
	if res.Error != nil {
		return comment, http.StatusInternalServerError, res.Error
	}
	return comment, http.StatusOK, nil
}
