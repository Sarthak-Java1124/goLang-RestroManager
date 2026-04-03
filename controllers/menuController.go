package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var menuCollection = database.OpenCollection(*database.DBinstance(), "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		page := c.Query("page")

		limit := c.Query("limit")

		matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
		facetStage := bson.D{{Key: "$facet", Value: bson.D{
			{Key: "metadata", Value: bson.D{{Key: "$count", Value: "total_count"}}},
			{Key: "data", Value: bson.A{
				bson.D{{Key: "$skip", Value: page}},
				bson.D{{Key: "$limit", Value: limit}},
			}},
		}}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "metadata", Value: bson.D{{Key: "$count", Value: "total_count"}}}, {Key: "data", Value: "$data"}}}}
		result, err := menuCollection.Aggregate(ctx, mongo.Pipeline{matchStage,
			facetStage,
			projectStage})
		if err != nil {
			log.Fatal("The error in menu aggregation is :", err)

		}
		var final_result bson.M

		if err = result.All(ctx, &final_result); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, final_result)

	}
}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("menu_id")
		var menu models.Menu
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": id}).Decode(&menu)
		if err != nil {
			log.Fatal("The error while fetching the menu by the id is : ", err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Successful get called ", "data": menu})

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menuBody models.Menu
		var foodBody models.FoodModel
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := c.BindJSON(&menuBody); err != nil {
			log.Fatal("The error in binding the request json in the menu controller is : ", err)
		}
		validationErr := validate.Struct(&menuBody)
		if validationErr != nil {
			fmt.Println("The validation error is : ", validationErr)
		}
		err := foodCollection.FindOne(ctx, bson.M{"food_id": menuBody.Food_id}).Decode(&foodBody)
		if err != nil {
			log.Fatal("The food id is invalid for the given menu")
		}

		menuBody.ID = primitive.NewObjectID()
		now := time.Now()
		menuBody.Created_at = now
		menuBody.Updated_at = now
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
