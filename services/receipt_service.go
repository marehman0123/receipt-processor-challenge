package services

import (
	"errors"
	"fmt"
	"math"
	"receipt-processor-challenge/models"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateReceipt(retailer, date, time, total string, items []models.Item) (models.Receipt, error) {
	receipt := models.Receipt{
		ID:           uuid.New().String(), // Generate a new UUID for each receipt
		Retailer:     retailer,
		PurchaseDate: date,
		PurchaseTime: time,
		Total:        total,
		Items:        items,
	}
	models.Receipts = append(models.Receipts, receipt)
	return receipt, nil
}

func GetReceiptByID(id string) (models.Receipt, error) {
	for _, receipt := range models.Receipts {
		if receipt.ID == id {
			return receipt, nil
		}
	}
	return models.Receipt{}, errors.New("receipt not found")
}

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	alphanumericCount := 0
	for _, char := range receipt.Retailer {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
			alphanumericCount++
		}
	}
	points += alphanumericCount

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil && total == float64(int(total)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil && math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points based on item description length being a multiple of 3
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		length := len(trimmedDescription)

		if length%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
				fmt.Println("PointsDescription:", points)
			}
		}
	}
	// Rule 6: 6 points if the day in the purchase date is odd
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && date.Day()%2 != 0 {
		points += 6
	}
	// Rule 7: 10 points if the purchase time is between 2:00 pm and 4:00 pm
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() == 14 || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
		points += 10
	}

	return points
}
