package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var userCollection = database.OpenCollection(*database.DBinstance(), "user")
var validate = validator.New()

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := userCollection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
		if err != nil {
			log.Fatal("The error in finding the user in the database is : ", err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Successfully fetched user", "data": user})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user LoginBody
		var userBody models.User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			log.Fatal("Error in binding the user json in login function is : ", err)
		}
		validationErr := validate.Struct(&user)
		if validationErr != nil {
			log.Fatal("The error in validating the login json is :", validationErr)
		}
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&userBody)
		if err != nil {
			log.Fatal("There is no user found in the collection during the login : ", err)
		}

	}
}
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func HashPassword(password string) string {

}

func VerifyHashPassword(password string, hashedPassword string) (bool, string) {

}
