package controllers

import (
	"api/config"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModels(c *gin.Context) {
	var models []models.Model
	config.DB.Find(&models)
	c.JSON(http.StatusOK, models)
}
func GetModel(c *gin.Context) {
	var models models.Model
	if err := config.DB.First(&models, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	c.JSON(http.StatusOK, models)
}

func CreateModel(c *gin.Context) {
	var models models.Model
	if err := c.ShouldBindJSON(&models); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&models)
	c.JSON(http.StatusCreated, models)
}

func UpdateModel(c *gin.Context) {
	var models models.Model
	if err := config.DB.First(&models, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	if err := c.ShouldBindJSON(&models); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&models)
	c.JSON(http.StatusOK, models)
}

func DeleteModel(c *gin.Context) {
	var models models.Model
	if err := config.DB.First(&models, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	config.DB.Delete(&models)
	c.JSON(http.StatusOK, gin.H{"message": "Model deleted"})
}
