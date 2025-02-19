package controllers

import (
	"api/config"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConditions(c *gin.Context) {
	var conditions []models.Condition
	config.DB.Find(&conditions)
	c.JSON(http.StatusOK, conditions)
}
func GetCondition(c *gin.Context) {
	var conditions models.Condition
	if err := config.DB.First(&conditions, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Condition not found"})
		return
	}
	c.JSON(http.StatusOK, conditions)
}

func CreateCondition(c *gin.Context) {
	var conditions models.Condition
	if err := c.ShouldBindJSON(&conditions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&conditions)
	c.JSON(http.StatusCreated, conditions)
}

func UpdateCondition(c *gin.Context) {
	var conditions models.Condition
	if err := config.DB.First(&conditions, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Condition not found"})
		return
	}
	if err := c.ShouldBindJSON(&conditions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&conditions)
	c.JSON(http.StatusOK, conditions)
}

func DeleteCondition(c *gin.Context) {
	var models models.Condition
	if err := config.DB.First(&models, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Condition not found"})
		return
	}
	config.DB.Delete(&models)
	c.JSON(http.StatusOK, gin.H{"message": "Condition deleted"})
}
