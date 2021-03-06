package services

import (
	"eden/models"
	"fmt"
	"strconv"
)

func QueryArticleList(pageStr string, sortStr string) (list []models.ExtendArticle, err error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	sort, err := strconv.Atoi(sortStr)
	if err != nil {
		return nil, err
	}
	list, err = models.QueryArticleList(page, sort)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func QueryArticle(idStr string) (art *models.ExtendArticle, err error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	art, err = models.QueryExtendArticle(id)
	if err != nil {
		return nil, err
	}
	return art, nil
}

func UpdateArticle(art *models.Article, u *models.User) (err error) {
	originArt := models.Article{ID: art.ID}
	err = models.QueryArticle(&originArt)
	if err != nil {
		return err
	}
	if originArt.CreatedBy != u.ID && !u.IsAdmin() {
		return fmt.Errorf("no access to this article")
	}
	return models.UpdateArticle(art)
}

func DeleteArticle(idStr string, u *models.User) (err error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	if !u.IsAdmin() {
		return fmt.Errorf("no access to delete this article")
	}

	return models.DeleteArticle(id)
}
