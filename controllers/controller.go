package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func SignUp(c echo.Context) error {
	msg := "signup controller works"

	return c.JSON(http.StatusOK, msg)
}
func LogIn(c echo.Context) error {
	msg := "login controller works"

	return c.JSON(http.StatusOK, msg)
}
