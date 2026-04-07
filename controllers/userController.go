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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = database.OpenCollection(*database.DBinstance(), "user")
var validate = validator.New()

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignUpBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
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
		var userBody SignUpBody
		if err := c.BindJSON(&userBody); err != nil {
			log.Fatal("The error in binding json in the signup body is : ", err)
		}
		validationErr := validate.Struct(&userBody)
		if validationErr != nil {
			log.Fatal("There is an error in validating the user body in the signup controller", validationErr)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var userModel models.User
		userModel.Email = &userBody.Email
		userModel.First_name = &userBody.First_name
		userModel.Last_name = &userBody.Last_name
		hashedPassword := HashPassword(userBody.Password)
		userModel.Password = &hashedPassword
		userModel.ID = primitive.NewObjectID()
		userModel.User_id = userModel.ID.Hex()
		now := time.Now()
		userModel.Created_at = now
		userModel.Updated_at = now
		refreshToken, err := utils.GenerateRefreshTokens()
		if err != nil {
			log.Fatal("The error in generating the refresh token in the signup controller is :", err)
		}
		hashedRefreshToken := utils.HashRefreshToken(refreshToken)
		userModel.Refresh_Token = &hashedRefreshToken
		user, err := userCollection.InsertOne(ctx, userModel)
		if err != nil {
			log.Fatal("The error in inserting user into mongo db is : ", err)
		}
		access_token := utils.GenerateJWTToken(userModel.ID, *userModel.Email)
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "data": user, "access_token": access_token})
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
		compare := VerifyHashPassword(user.Password, *userBody.Password)
		if compare != nil {
			log.Fatal("The password doesn't match")
			return
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
		var userBody models.User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		data, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal("The error in finding user is : ", err)
		}
		data.All(ctx, userBody)

		c.JSON(http.StatusOK, gin.H{"message": "Success returning users", "data": data})
	}
}
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("The error in hashing password from the given password is : ", err)
	}
	return string(hashedPassword)
}

func VerifyHashPassword(password string, hashedPassword string) error {
	compare := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return compare
}
