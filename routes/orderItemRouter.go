package routes

import (
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItem", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItem/:orderTtem_id", controllers.GetMenu())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrder())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())

}
