package services

import (
    "fmt"
	"eden/models"
)

func ArtList(page int, sort int) (list []models.ExtArt, err error) {
    list, err = models.ArtListQuery(page, sort)
    if err != nil {
        return nil, err
    }
	return list, nil
}

func ShowArt(id uint64) (art *models.ExtArt, err error) {
	return nil, nil
}

func UpdateArt(id uint64, art *models.Article, u *models.User) (err error) {
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

func DeleteArt(id uint64, u *models.User) (err error) {
    return nil
}