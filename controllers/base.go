package controllers

import (
    "eden/models"
    "eden/config"
    
	"html/template"
	"net/http"
    "time"
    "io"

	"github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
)

var T *Template

type Template struct {
    templates *template.Template
}

func About(c echo.Context) error {
	return c.Render(http.StatusOK, "about", nil)
}

func NotFound(c echo.Context) error {
    return c.Render(http.StatusNotFound, "404", nil)
}

func Friendship(c echo.Context) error {
	return c.Render(http.StatusOK, "friendship", nil)
}

//jwt auth
func SetUser(user *models.User, c echo.Context) error {
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims["user"] = user
    token.Claims["exp"] = time.Now().Add(time.Hour * 720).Unix()
    tokenString, err := token.SignedString([]byte(config.Config.JwtAuthKey))
	if err != nil {
		return err
	}
    c.Response().Header().Set("Authorization", "bear " + tokenString)
    return nil
}

func currentUser(c *echo.Context) *models.User {
    token := c.Get("user").(*jwt.Token)
    user, ok := token.Claims["user"].(models.User)
    if !ok {
        return nil
    }
	return &user
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func init() {
    T = &Template{
    	templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}