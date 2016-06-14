package controllers

import (
	"eden/config"
	"eden/models"

	"html/template"
	"io"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

var T *Template

type Template struct {
	templates *template.Template
}

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/login"
}

func (Auth) LogoutURL(*admin.Context) string {
	return "/logout"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, err){
		
	})
	user, ok := token.Claims["user"].(models.User)
	if !ok {
		return nil
	}
	if !user.IsAdmin() {
		return nil
	}
	return &user
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
	c.Response().Header().Set("Authorization", tokenString)
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
