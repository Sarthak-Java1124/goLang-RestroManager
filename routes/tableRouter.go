package routes

import (
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())

}
