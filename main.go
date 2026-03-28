package main

import (
	"os"

	"github.com/Sarthak-Java1124/goLang-RestroManager.git/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger)
	routes.UserRoutes(router)
	router.Use(middleware.Authentication)
	routes.FoodRoutes(router)
	routes.TableRoute(router)
	routes.OrderRoute(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)

}
