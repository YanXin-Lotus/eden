package controllers

import "net/http"

type MainController struct {
	BaseController
}

func (c MainController) Index() error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func (c MainController) Pagination() error {
	return c.String(http.StatusOK, "Pagination\n")
}

func (c MainController) Category() error {
	return c.Stirng(http.StatusOK, "Category\n")
}
