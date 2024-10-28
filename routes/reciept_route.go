package routes

import (
	"receipt-processor-challenge/controllers"

	"receipt-processor-challenge/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	receiptRoutes := router.Group("/receipts")
	{
		receiptRoutes.GET("/:id/points", middlewares.ValidateUUIDMiddleware(), controllers.CalculateReceiptPoints)
		receiptRoutes.POST("/process", middlewares.ValidateReceiptMiddleware(), controllers.CreateReceipt)
	}
}
