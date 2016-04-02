package controllers

import (
    "eden/models"
    "eden/config"
    
	"html/template"
	"net/http"
    "time"
    "fmt"

	"github.com/gocraft/web"
	"github.com/unrolled/render"
    "github.com/dgrijalva/jwt-go"
)

var (
    Render *render.Render
    BaseContext *Context
)

type Context struct {
    *render.Render
    models.User
}

func (c *Context) Prepare(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	return
}

func (c *Context) About(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "about", nil)
}

func (c *Context) NotFound(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    c.HTML(rw, http.StatusNotFound, "404", nil)
}

func (c *Context) Friendship(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "friendship", nil)

}

func (c *Context) Redirect(rw web.ResponseWriter, req *web.Request, path string) {
    http.Redirect(rw, req.Request, path, http.StatusFound)
}

//jwt auth
func (c *Context) SetUser(user *models.User, rw http.ResponseWriter) error {
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims["user"] = user
    token.Claims["exp"] = time.Now().Add(time.Hour * 720).Unix()
    tokenString, err := token.SignedString([]byte(config.Config.JwtAuthKey))
	if err != nil {
		return err
	}
    rw.Header().Set("Authorization", "bear " + tokenString)
    return nil
}

func (c *Context) CurrentUser(req *web.Request) *models.User {
    token, err := jwt.ParseFromRequest(req.Request, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.Config.JwtAuthKey), nil
    })
    
    if err != nil {
        fmt.Println("parse token err :", err)
        return nil
    }
    
    user := token.Header["user"].(models.User)
    expire := token.Header["exp"].(string)
    fmt.Println("expire is :", expire)
    
	return &user
}

func init() {
	Render = render.New(render.Options{
		Directory:  "public/views",
		Extensions: []string{".html"},
		Funcs:      []template.FuncMap{},
	})
    
    BaseContext.Render = Render
}
