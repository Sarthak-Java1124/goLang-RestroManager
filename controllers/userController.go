package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/models"
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/utils"
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

		refresh_token, err := utils.GenerateRefreshTokens()
		if err != nil {
			log.Fatal("The error in generating refresh_token", err)
		}
		tokenHash := utils.HashRefreshToken(refresh_token)
		_, _ = userCollection.UpdateOne(
			ctx,
			bson.M{"_id": userBody.ID},
			bson.M{"$set": bson.M{"refresh_token": tokenHash}},
		)
		access_token := utils.GenerateJWTToken(userBody.ID, *userBody.Email)
		c.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refresh_token,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
		})
		userBody.Refresh_Token = &tokenHash
		c.JSON(http.StatusOK, gin.H{"message": "Logged in Successfully", "access_token": access_token})

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
