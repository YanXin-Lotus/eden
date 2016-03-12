package controllers

import "net/http"

type MainController struct {
	BaseController
}

func (c MainController) Index() error {
	return c.Context.String(http.StatusOK, "Hello, World!\n")
}

func (c MainController) Pagination() error {
	return c.Context.String(http.StatusOK, "Pagination\n")
}

func (c MainController) Category() error {
	return c.Context.String(http.StatusOK, "Category\n")
}
