package services

import (
	"eden/models"
	"fmt"
	"strconv"
)

func ArtList(pageStr string, sortStr string) (list []models.ExtArt, err error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	sort, err := strconv.Atoi(sortStr)
	if err != nil {
		return nil, err
	}
	list, err = models.ArtListQuery(page, sort)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func QueryArt(idStr string) (art *models.ExtArt, err error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	art, err = models.ShowArt(id)
	if err != nil {
		return nil, err
	}
	return art, nil
}

func UpdateArt(idStr string, art *models.Article, u *models.User) (err error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	originArt := models.Article{ID: id}
	err = models.ArtQuery(&originArt)
	if err != nil {
		return err
	}
	if originArt.CreatedBy != u.ID && !u.IsAdmin() {
		return fmt.Errorf("no access to this article")
	}
	return models.UpdateArt(art)
}

func DeleteArt(idStr string, u *models.User) (err error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	if !u.IsAdmin() {
		return fmt.Errorf("no access to delete this article")
	}

	return models.DeleteArt(id)
}
