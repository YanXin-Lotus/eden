package controllers

import (
	"eden/models"

	"encoding/gob"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

var T *Template
var SessionStore *sessions.CookieStore

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
	session, _ := SessionStore.Get(c.Request, "eden")
	suser := session.Values["user"]
	if suser == nil {
		return nil
	}
	user := suser.(*models.User)
	if !user.IsAdmin() {
		return nil
	}
	return user
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

//set session
func setUser(user *models.User, c echo.Context) error {
	session, _ := SessionStore.Get(c.Request(), "eden")
	session.Values["user"] = user
	session.Save(c.Request(), c.Response().Writer)
	return nil
}

func currentUser(c echo.Context) *models.User {
	session, _ := SessionStore.Get(c.Request(), "eden")
	user := session.Values["user"]
	if user == nil {
		return nil
	}
	return user.(*models.User)
}

func delSession(c echo.Context) error {
	session, _ := SessionStore.Get(c.Request(), "eden")
	session.Values["user"] = nil
	return session.Save(c.Request(), c.Response().Writer)
}

// Render render a template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {
	gob.Register(&models.User{})
	SessionStore = sessions.NewCookieStore([]byte("I'mSecretKey"))
	T = &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}
