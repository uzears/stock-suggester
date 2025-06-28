package controllers

import (
	"net/http"
	"stock-suggester/services"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func GetNewsSuggestions(c *gin.Context) {
	suggestions, err := services.GenerateSuggestions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, suggestions)
}
