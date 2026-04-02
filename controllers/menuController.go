package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var menuCollection = database.OpenCollection(*database.DBinstance(), "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
     page := c.Query("page")
	 limit := c.Query("limit")
     
	 matchStage := bson.D{{Key: "$match" , Value: bson.D{}}}
	facetStage := bson.D{{Key: "$facet", Value: bson.D{
			{Key: "metadata", Value: bson.D{{Key: "$count", Value: "total_count"}}},
			{Key: "data", Value: bson.A{
				bson.D{{Key: "$skip", Value: startIndex}},
				bson.D{{Key: "$limit", Value: endIndex}},
			}},
		}}}
	projectStage := bson.D{{Key:  "$project" , Value: bson.D{{Key : "metadata"  , Value : bson.D{{Key : "$count" , Value:  "total_count"}}} , {Key: "data" , Value : "$data"}}}}
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

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
