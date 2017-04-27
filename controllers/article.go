package controllers

import (
	"eden/models"
	"eden/services"
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	data := make(map[string]interface{})
	list, err := services.QueryArticleList("1", "0")
	if err != nil {
		return c.Redirect(http.StatusOK, "/404")
	}
	data["list"] = list
	data["user"] = currentUser(c)
	return c.Render(http.StatusOK, "index", data)
}

func Pagination(c echo.Context) error {
	data := make(map[string]interface{})
	page := c.Param("page")
	sort := c.Param("sort")
	list, err := services.QueryArticleList(page, sort)
	if err != nil {
		return c.Redirect(http.StatusOK, "/404")
	}
	data["list"] = list
	data["user"] = currentUser(c)
	return c.Render(http.StatusOK, "pagination", data)
}

func Category(c echo.Context) error {
	data := make(map[string]interface{})
	page := c.Param("page")
	sort := c.Param("sort")
	list, err := services.QueryArticleList(page, sort)
	if err != nil {
		return c.Redirect(http.StatusOK, "/")
	}
	data["list"] = list
	data["user"] = currentUser(c)
	return c.Render(http.StatusOK, "pagination", data)
}

func Article(c echo.Context) error {
	data := make(map[string]interface{})
	id := c.Param("id")
	art, err := services.QueryArticle(id)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/404")
	}
	data["art"] = art
	data["user"] = "currentuser"
	return c.Render(http.StatusOK, "detail", data)
}

func EditArticle(c echo.Context) error {
	data := make(map[string]interface{})
	id := c.Param("id")
	art, err := services.QueryArticle(id)
	if err != nil {
		return c.Redirect(http.StatusOK, "/404")
	}
	data["art"] = art
	data["user"] = currentUser(c)
	return c.Render(http.StatusOK, "edit", data)
}

func DoEditArticle(c echo.Context) error {
	data := make(map[string]interface{})
	user := currentUser(c)
	var art models.Article
	err := c.Bind(art)
	if err != nil {
		return c.JSON(http.StatusOK, &retJson{OK: false, Desc: "err bind article"})
	}
	err = services.UpdateArticle(&art, user)
	if err != nil {
		return c.JSON(http.StatusOK, &retJson{OK: false, Desc: "err update article"})
	}
	return c.Render(http.StatusOK, "detail", data)
}
