package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(*database.DBinstance(), "food")
var validate = validator.New()

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
				bson.D{{Key: "$limit", Value: endIndex}},
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
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing food items"})
		}
		var allFoods []bson.M
		if err = result.All(ctx, &allFoods); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allFoods)

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foodBody models.FoodModel
		var menuBody models.Menu
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		if err := c.Bind(&foodBody); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "An error occured while binding"})
		}
		validationErr := validate.Struct(&foodBody)
		if validationErr != nil {
			fmt.Println("The validation error is : ", validationErr)
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": foodBody.Menu_id}).Decode(&menuBody)
		if err != nil {
			fmt.Println("Invalid Menu Id for the menu collection", err)
		}
		foodBody.ID = primitive.NewObjectID()
		now := time.Now()
		foodBody.Created_at = &now
		foodBody.Updated_at = &now
		foodBody.Food_id = foodBody.ID.Hex()
		var num = toFixed(*foodBody.Price, 2)
		foodBody.Price = &num

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

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
