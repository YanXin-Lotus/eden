package routers

import (
    "github.com/labstack/echo"
    "eden/controllers"
)

func index() echo.HandlerFunc {
    return controllers.Index
}

func about() echo.HandlerFunc {
    return controllers.About
}

func friendship() echo.HandlerFunc {
    return controllers.Friendship
}

func pagination() echo.HandlerFunc {
    return controllers.Pagination
}

func category() echo.HandlerFunc {
    return controllers.Category
}

func article() echo.HandlerFunc {
    return controllers.Article
}

func editArticle() echo.HandlerFunc {
    return controllers.EditArticle
}

func doEditArticle() echo.HandlerFunc {
    return controllers.DoEditArticle
}

func login() echo.HandlerFunc {
    return controllers.Login
}

func doLogin() echo.HandlerFunc {
    return controllers.DoLogin
}

func signout() echo.HandlerFunc {
    return controllers.Signout
}

func register() echo.HandlerFunc {
    return controllers.Register
}

func doRegister() echo.HandlerFunc {
    return controllers.DoRegister
}

func info() echo.HandlerFunc {
    return controllers.Info
}

func editInfo() echo.HandlerFunc {
    return controllers.EditInfo
}

func doEditInfo() echo.HandlerFunc {
    return controllers.DoEditInfo
}

func changePW() echo.HandlerFunc {
    return controllers.ChangePW
}

func doChangePW() echo.HandlerFunc {
    return controllers.DoChangePW
}