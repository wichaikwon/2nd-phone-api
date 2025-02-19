package controllers

import (
	"api/config"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPricingAdjustments(c *gin.Context) {
	var pricing_adjustment []models.PricingAdjustment
	config.DB.Find(&pricing_adjustment)
	c.JSON(http.StatusOK, pricing_adjustment)
}

func GetPricingAdjustment(c *gin.Context) {
	var pricing_adjustment models.PricingAdjustment
	if err := config.DB.First(&pricing_adjustment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PricingAdjustment not found"})
		return
	}
	c.JSON(http.StatusOK, pricing_adjustment)
}

func CreatePricingAdjustment(c *gin.Context) {
	var pricing_adjustment models.PricingAdjustment
	if err := c.ShouldBindJSON(&pricing_adjustment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&pricing_adjustment)
	c.JSON(http.StatusCreated, pricing_adjustment)
}

func UpdatePricingAdjustment(c *gin.Context) {
	var pricing_adjustment models.PricingAdjustment
	if err := config.DB.First(&pricing_adjustment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PricingAdjustment not found"})
		return
	}
	if err := c.ShouldBindJSON(&pricing_adjustment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&pricing_adjustment)
	c.JSON(http.StatusOK, pricing_adjustment)
}

func DeletePricingAdjustment(c *gin.Context) {
	var pricing_adjustment models.PricingAdjustment
	if err := config.DB.First(&pricing_adjustment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PricingAdjustment not found"})
		return
	}
	config.DB.Delete(&pricing_adjustment)
	c.JSON(http.StatusOK, gin.H{"message": "PricingAdjustment deleted"})
}
