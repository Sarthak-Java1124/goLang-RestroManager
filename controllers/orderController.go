package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Sarthak-Java1124/goLang-RestroManager.git/database"
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var orderCollection = database.OpenCollection(*database.DBinstance(), "order")
var validate = validator.New()

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		page := c.Query("page")
		startIndex, err := strconv.Atoi(page)
		if err != nil {
			log.Fatal("The error in converting the page to int is : ", err)

		}
		limit := c.Query("limit")
		limitIndex, err := strconv.Atoi(limit)
		if err != nil {
			log.Fatal("The error in converting the page to int is : ", err)
		}
		matchStage := bson.D{{Key: "$match", Value: bson.M{}}}
		facetStage := bson.D{{Key: "$facet", Value: bson.D{{Key: "metadata", Value: bson.D{{Key: "$count", Value: "total_count"}}}, {Key: "data", Value: bson.D{{Key: "$skip", Value: startIndex}, {Key: "$limit", Value: limitIndex}}}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"total_count", bson.D{
					{"$ifNull", bson.A{
						bson.D{{"$arrayElemAt", bson.A{"$metadata.total_count", 0}}},
						0,
					}},
				}},
				{"order_items", "$data"},
			}},
		}
		result, err := orderCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			facetStage,
			projectStage,
		})
		if err != nil {
			log.Fatal("The error in inserting the result in order collection is : ", err)
		}
		var final_result bson.M
		result.All(ctx, &final_result)
		c.JSON(http.StatusOK, gin.H{"message": "Result fetched success", "data": final_result})

	}

}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("order_id")
		var orderBody models.Order
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := orderCollection.FindOne(ctx, bson.M{"order_id": id}).Decode(&orderBody)
		if err != nil {
			log.Fatal("The was a problem in decoding the order into the order body")

		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"message": "Succesfully fetched the orders", "data": orderBody})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		var table models.Table
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		validationErr := validate.Struct(&order)
		if validationErr != nil {
			log.Fatal("The error in validating the request body is : ", err)
		}
		if err := c.BindJSON(&order); err != nil {
			log.Fatal("The error in binding the json in the order controller is :", err)
		}
		err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
		if err != nil {
			log.Fatal("The error in finding the table id for the order controller is : ", err)
		}

		order.ID = primitive.NewObjectID()
		now := time.Now()
		order.Created_at = now
		order.Updated_at = now
		order.Order_id = order.ID.Hex()
		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			msg := fmt.Sprintf("order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Order sent successfully", "data": result})

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
