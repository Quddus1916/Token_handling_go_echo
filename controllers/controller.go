package controllers

import (
	"github.com/Quddus1916/Token_handling_go_echo/database"
	"github.com/Quddus1916/Token_handling_go_echo/helpers"
	"github.com/Quddus1916/Token_handling_go_echo/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"golang.org/x/crypto/bcrypt"
	// /"github.com/labstack/echo"
	"context"
	//"log"
	"net/http"
)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func SignUp(c echo.Context) error {
	//	get user data
	var userData models.User
	if err := c.Bind(userData); err != nil {
		msg := "failed to bind data"
		return c.JSON(http.StatusBadRequest, msg)
	}

	//validate user data
	if validationerr := validate.Struct(userData); validationerr != nil {
		msg := "data is not valid"
		return c.JSON(http.StatusBadRequest, msg)
	}
	// validate user email and phone number
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	countemail, _ := UserCollection.CountDocuments(ctx, bson.M{"email": userData.Email})
	defer cancel()
	countphone, _ := UserCollection.CountDocuments(ctx, bson.M{"phone": userData.Phone})
	defer cancel()
	if countemail > 0 || countphone > 0 {
		msg := "email or phone number already exist"
		return c.JSON(http.StatusBadRequest, msg)
	}

	userData.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	userData.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	userData.ID = primitive.NewObjectID()
	userData.User_id = userData.ID.Hex()
	token, refresh_token, err := helpers.GenerateAllTokens(*userData.Email, *userData.First_name, *userData.Last_name, *userData.User_type, *&userData.User_id)
	userData.Token = &token
	userData.Refresh_token = &refresh_token
	result, err := UserCollection.InsertOne(ctx, userData)
	if err != nil {
		msg := "data can't be inserted"
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.JSON(http.StatusOK, result)
}
func LogIn(c echo.Context) error {
	msg := "login controller works"

	return c.JSON(http.StatusOK, msg)
}
func GetUsers(c echo.Context) error {

	msg := "login controller works"
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
