package main

import (
	"github.com/gin-gonic/gin"

	"receipt-processor-challenge/routes"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(router)
	router.Run(":8080")
}
