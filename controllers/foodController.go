package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(*database.DBinstance(), "food")

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		foodId := c.Param("food_id")
		var food models.FoodModel
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		jsonData, err := json.Marshal(food)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "An error occurred while getting food"})
		}
		c.JSON(http.StatusOK, gin.H{"data": jsonData})

	}
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Fatal(err)
		}
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			log.Fatal(err)
		}

		startIndex := page
		endIndex := limit

		matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
		facetStage := bson.D{{Key: "$facet", Value: bson.D{
			{Key: "metadata", Value: bson.D{{Key: "$count", Value: "total_count"}}},
			{Key: "data", Value: bson.A{
				bson.D{{Key: "$skip", Value: startIndex}},
				bson.D{{Key: "$limit", Value: limit}},
			}},
		}}}

	
	projectStage := bson.D{
		{"$project", bson.D{
			{"total_count", bson.D{
				{"$ifNull", bson.A{
					bson.D{{"$arrayElemAt", bson.A{"$metadata.total_count", 0}}},
					0,
				}},
			}},
			{"food_items", "$data"},
		}},
	}
	result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		facetStage,
		projectStage,
	})
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foodBody models.FoodModel
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		if err := c.Bind(&foodBody); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "An error occured while binding"})
		}
		data, err := foodCollection.InsertOne(ctx, foodBody)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "An error occured while inserting the doc"})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Successfully entered the data", "data": data})

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) {

}

func toFixed(num float64, precision int) {

}
