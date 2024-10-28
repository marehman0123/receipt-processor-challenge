package middlewares

import (
	"net/http"
	"regexp"
	"strings"

	"receipt-processor-challenge/models" // replace with your actual models path

	"github.com/gin-gonic/gin"
)

var (
	retailerPattern         = regexp.MustCompile(`^[\w\s\-&]+$`)
	datePattern             = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	timePattern             = regexp.MustCompile(`^\d{2}:\d{2}$`)
	totalPattern            = regexp.MustCompile(`^\d+\.\d{2}$`)
	shortDescriptionPattern = regexp.MustCompile(`^[\w\s\-]+$`)
	pricePattern            = regexp.MustCompile(`^\d+\.\d{2}$`)
)

// ValidateReceiptMiddleware validates the input for a receipt based on predefined patterns
func ValidateReceiptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var receipt models.Receipt
		if err := c.ShouldBindJSON(&receipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			c.Abort()
			return
		}

		// Validate retailer
		if !retailerPattern.MatchString(receipt.Retailer) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid retailer format"})
			c.Abort()
			return
		}

		// Validate purchaseDate
		if !datePattern.MatchString(receipt.PurchaseDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchaseDate format"})
			c.Abort()
			return
		}

		// Validate purchaseTime
		if !timePattern.MatchString(receipt.PurchaseTime) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchaseTime format"})
			c.Abort()
			return
		}

		// Validate total
		if !totalPattern.MatchString(receipt.Total) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid total format"})
			c.Abort()
			return
		}

		// Validate items
		if len(receipt.Items) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "At least one item is required"})
			c.Abort()
			return
		}

		for _, item := range receipt.Items {
			if !shortDescriptionPattern.MatchString(strings.TrimSpace(item.ShortDescription)) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item shortDescription format"})
				c.Abort()
				return
			}
			if !pricePattern.MatchString(item.Price) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item price format"})
				c.Abort()
				return
			}
		}

		// Store validated receipt in context
		c.Set("validatedReceipt", receipt)
		// If all validations pass, proceed to the next handler
		c.Next()
	}
}

// UUID regex pattern for validation (UUID version 4)
var uuidPattern = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

// ValidateUUIDMiddleware checks if the 'id' path parameter is a valid UUID
func ValidateUUIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id") // Retrieve the 'id' path parameter

		if !uuidPattern.MatchString(id) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			c.Abort()
			return
		}

		c.Next() // Continue if UUID is valid
	}
}
