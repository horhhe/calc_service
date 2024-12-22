package main

import (
	"github.com/gin-gonic/gin"
	"github.com/horhhe/calc_service/internal/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/api/v1/calculate", handlers.CalculateExpression)

	r.Run(":8080")
}
