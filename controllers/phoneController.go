package controllers

import (
	"api/config"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPhones(c *gin.Context) {
	var phones []models.Phone
	config.DB.Find(&phones)
	c.JSON(http.StatusOK, phones)
}

func GetPhone(c *gin.Context) {
	var phone models.Phone
	if err := config.DB.First(&phone, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	c.JSON(http.StatusOK, phone)
}

func CreatePhone(c *gin.Context) {
	var phone models.Phone
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&phone)
	c.JSON(http.StatusCreated, phone)
}

func UpdatePhone(c *gin.Context) {
	var phone models.Phone
	if err := config.DB.First(&phone, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&phone)
	c.JSON(http.StatusOK, phone)
}

func DeletePhone(c *gin.Context) {
	var phone models.Phone
	if err := config.DB.First(&phone, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	config.DB.Delete(&phone)
	c.JSON(http.StatusOK, gin.H{"message": "Phone deleted"})
}
