package controllers

import (
    "eden/services"
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
    data := make(map[string]interface{})
    list, err := services.ArtList("1", "0")
    if err != nil {
        return c.Redirect(http.StatusOK, "/404")
    }
    data["list"] = list
    data["user"] = "currentuser"
	return c.Render(http.StatusOK, "index", data)
}

func Pagination(c echo.Context) error {
    data := make(map[string]interface{})
    page := c.Param("page")
    sort := c.Param("sort")
    list, err := services.ArtList(page ,sort)
    if err != nil {
        return c.Redirect(http.StatusOK, "/404")
    }
    data["list"] = list
    data["user"] = "currentuser"
	return c.Render(http.StatusOK, "pagination", data)
}

func Category(c echo.Context) error {
    data := make(map[string]interface{})
    page := c.Param("page")
    sort := c.Param("sort")
    list, err := services.ArtList(page, sort)
    if err != nil {
        return c.Redirect(http.StatusOK, "/")
    }
	data["list"] = list
    data["user"] = "currentuser"
	return c.Render(http.StatusOK, "pagination", data)
}

func Article(c echo.Context) error {
    data := make(map[string]interface{})
    id := c.Param("id")
    art, err := services.ShowArt(id)
    if err != nil {
        return c.Redirect(http.StatusTemporaryRedirect, "/404")
    }
	data["art"] = art
    data["user"] = "currentuser"
	return c.Render(http.StatusOK, "detail", data)
}

func EditArticle(c echo.Context) error {
	return c.Render(http.StatusOK, "edit", nil)
}

func DoEditArticle(c echo.Context) error {
    user := currentUser(c)
    data := make(map[string]interface{})
    id := c.Param("id")
	err := services.UpdateArt(id, nil, user)
    if err != nil {
        return c.Redirect(http.StatusTemporaryRedirect, "/404")
    }
    return c.Render(http.StatusOK, "detail", data)
}
