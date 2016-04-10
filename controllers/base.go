package controllers

import (
    "eden/models"
    "eden/config"
    
    "strconv"
	"html/template"
	"net/http"
    "time"
    "fmt"

	"github.com/gocraft/web"
	"github.com/unrolled/render"
    "github.com/dgrijalva/jwt-go"
    "github.com/mitchellh/mapstructure"
)

var (
    Render *render.Render
    BaseContext *Context
)

type Context struct {
    *render.Render
    models.User
    Data map[string]interface{}
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

func (c *Context) ParseParam2Int(req *web.Request, name string) int {
    dataString := req.PathParams[name]
    data, err := strconv.Atoi(dataString)
    if err != nil {
        panic(err)
    }
    return data
}

func (c *Context) ParseParam2Uint64(req *web.Request, name string) uint64 {
    datastring := req.PathParams[name]
    data, err := strconv.ParseUint(datastring, 10, 64)
    if err != nil {
        panic(err)
    }
    return data
}

func (c *Context) Parse2Struct(req *web.Request, obj interface{}) (err error) {
    err = mapstructure.Decode(req.Form, obj)
    return err
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
