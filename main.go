package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CalculateInput struct {
	Data []float64 `json:"data" binding:"required,min=1,dive,gt=0.0"`
}

type CalculateRes struct {
	Total float64 `json:"total"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Avg   float64 `json:"avg"`
}

type HealthRes struct {
	Status string `json:"status"`
}

func main() {
	r := gin.Default()
	r.GET("/health", GetHealthHandler)
	r.POST("/calculate", PostCalculateHandler)

    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "3000"
    }

    err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Impossible to start server: %s", err)
	}
}

func GetHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthRes{Status: "OK"})
}

func PostCalculateHandler(c *gin.Context) {
	var request CalculateInput

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := CalculateRes{
		Total: 0.0,
		Min:   0.0,
		Max:   0.0,
		Avg:   0.0,
	}

	for _, v := range request.Data {
		res.Total += v
		if res.Min == 0.0 || v < res.Min {
			res.Min = v
		}
		if v > res.Max {
			res.Max = v
		}
	}

	res.Avg = roundNum(res.Total / float64(len(request.Data)))

	res.Total = roundNum(res.Total)

	c.JSON(http.StatusOK, res)
}

func roundNum(num float64) float64 {
	return math.Round(num*100) / 100
}
