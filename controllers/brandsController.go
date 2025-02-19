package controllers

import (
	"api/config"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBrands(c *gin.Context) {
	var brands []models.Brand
	config.DB.Find(&brands)
	c.JSON(http.StatusOK, brands)
}
func GetBrand(c *gin.Context) {
	var brand models.Brand
	if err := config.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func CreateBrand(c *gin.Context) {
	var brand models.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&brand)
	c.JSON(http.StatusCreated, brand)
}

func UpdateBrand(c *gin.Context) {
	var brand models.Brand
	if err := config.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&brand)
	c.JSON(http.StatusOK, brand)
}

func DeleteBrand(c *gin.Context) {
	var brand models.Brand
	if err := config.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	config.DB.Delete(&brand)
	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted"})
}
