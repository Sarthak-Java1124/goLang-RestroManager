package routes

import "github.com/gin-gonic/gin"

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItem", controller.GetOrderItem())
	incomingRoutes.GET("/orderItem/:orderTtem_id", controller.GetMenu())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrder())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

}
