package routes

import (
	"github.com/Sarthak-Java1124/goLang-RestroManager.git/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoices())
	incomingRoutes.POST("/invoices", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())

}
