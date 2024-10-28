package controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"receipt-processor-challenge/models"
	"receipt-processor-challenge/services"
)

func CreateReceipt(c *gin.Context) {
	receiptInterface, _ := c.Get("validatedReceipt")
	newReceipt := receiptInterface.(models.Receipt)

	receipt, err := services.CreateReceipt(newReceipt.Retailer, newReceipt.PurchaseDate, newReceipt.PurchaseTime, newReceipt.Total, newReceipt.Items)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": receipt.ID})
}

func CalculateReceiptPoints(c *gin.Context) {

	id := c.Param("id")
	receipt, err := services.GetReceiptByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	points := services.CalculatePoints(receipt)
	c.JSON(http.StatusOK, gin.H{"points": points})
}
