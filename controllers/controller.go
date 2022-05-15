package controllers

import (
	"github.com/Quddus1916/Token_handling_go_echo/database"
	"github.com/Quddus1916/Token_handling_go_echo/helpers"
	"github.com/Quddus1916/Token_handling_go_echo/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"golang.org/x/crypto/bcrypt"
	// /"github.com/labstack/echo"
	"context"
	"net/http"
)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func SignUp(c echo.Context) error {
	msg := "signup controller works"

	return c.JSON(http.StatusOK, msg)
}
func LogIn(c echo.Context) error {
	msg := "login controller works"

	return c.JSON(http.StatusOK, msg)
}
func GetUsers(c echo.Context) error {
	msg := "get users  controller works"

	return c.JSON(http.StatusOK, msg)
}

func Getuser(c echo.Context) error {
	var user_id = new(models.Id)
	err := c.Bind(user_id)
	if err != nil {
		msg := "failed to bind data"
		return c.JSON(http.StatusOK, msg)
	}

	err = helpers.MatchUserTypeToUserId(c, user_id.ID)
	if err != nil {
		msg := "user type not matched"
		return c.JSON(http.StatusBadRequest, msg)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var user models.User
	err = UserCollection.FindOne(ctx, bson.M{"user_id": user_id.ID}).Decode(&user)
	defer cancel()
	if err != nil {
		//msg := "user data not available"
		return c.JSON(http.StatusBadRequest, user_id.ID)

	}

	return c.JSON(http.StatusOK, user)
}

func HashPassword()   {}
func VerifyPassword() {}
