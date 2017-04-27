package services

import (
	"eden/models"
	"eden/utils"
	"fmt"
)

func Login(u *models.User) (err error) {
	u.Password = utils.MD5(u.Password)
	err = models.QueryUser(u)
	return err
}

func Register(u *models.User) (err error) {
	if checkUserName(&models.User{Name: u.Name}) {
		return fmt.Errorf("username exsists!")
	}

	if checkEmail(&models.User{Email: u.Email}) {
		return fmt.Errorf("email exsists!")
	}
	return models.CreateUser(u)
}

func UserInfo(uid int) (u *models.User, err error) {
	user := models.User{ID: uid}
	err = models.QueryUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func EditUserInfo(uid int, u *models.User) (err error) {
	if uid != u.ID {
		return fmt.Errorf("no access to edit this user info")
	}

	return models.UpdateUser(u)
}

func checkUserName(u *models.User) bool {
	return models.UserExsist(u)
}

func checkEmail(u *models.User) bool {
	return models.UserExsist(u)
}
