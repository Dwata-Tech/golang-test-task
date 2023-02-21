package service

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/Dwata-Tech/golang-test-task/database"
	"github.com/Dwata-Tech/golang-test-task/model"
	"github.com/Dwata-Tech/golang-test-task/rabbitmq"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func CreateArticleService(article model.Article) (map[string]string, int, error) {
	var res = make(map[string]string)

	//request validation
	if len(article.Title) == 0 {
		return res, http.StatusBadRequest, errors.New("title can not be null or blank")
	}

	if len(article.Content) == 0 {
		return res, http.StatusBadRequest, errors.New("content can not be null or blank")
	}

	if len(article.Nickname) == 0 {
		return res, http.StatusBadRequest, errors.New("nickname can not be null or blank")
	}

	err := rabbitmq.PublishMessage(EncodeToBytes(article))
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	//res := database.Instance.Create(&article)
	//if res.Error != nil {
	//	return article, http.StatusInternalServerError, res.Error
	//}
	res["message"] = "Article pushed to queue successfully"
	return res, http.StatusOK, nil
}

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func GetArticleService(articleId int) (model.Article, int, error) {

	var article model.Article
	result := database.Instance.First(&article, articleId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return article, http.StatusBadRequest, errors.New("article not found")
	} else if result.Error != nil {
		logrus.Error("Error: " + result.Error.Error())
		return article, http.StatusInternalServerError, errors.New("internal server error")
	}

	return article, http.StatusOK, nil
}

func GetArticleListService(pageNumber, pageSize int) ([]model.ArticleListResponse, int, error) {

	// Calculate offset and limit
	offset := (pageNumber - 1) * pageSize
	limit := pageSize

	var response []model.ArticleListResponse
	var articles []model.Article
	result := database.Instance.Offset(offset).Limit(limit).Find(&articles)
	if result.Error != nil {
		return response, http.StatusInternalServerError, errors.New("internal server error")
	}
	for _, val := range articles {
		response = append(response, model.ArticleListResponse{
			Nickname:  val.Nickname,
			Title:     val.Title,
			CreatedAt: val.CreatedAt,
		})
	}
	return response, http.StatusOK, nil
}
