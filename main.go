package main

import (
	"github.com/Quddus1916/Token_handling_go_echo/controllers"

	"github.com/labstack/echo"

	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	e := echo.New()
	e.GET("/api-1", func(c echo.Context) error {
		msg := "access Granted for api-1"
		return c.JSON(http.StatusOK, msg)
	})

	e.GET("/api-2", func(c echo.Context) error {
		msg := "access Granted for api-2"
		return c.JSON(http.StatusOK, msg)
	})

	e.POST("/users/signup", controllers.SignUp)
	e.POST("/users/login", controllers.LogIn)
	e.GET("/user", controllers.Getuser)
	e.GET("/users", controllers.GetUsers)
	e.Start(":" + port)

}
