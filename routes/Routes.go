package routes

import (
	"api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// authMiddleware ตัวอย่าง middleware ที่ตรวจสอบ Authorization
func authMiddleware(c *gin.Context) {
	// ตัวอย่างการตรวจสอบ header Authorization
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "Authorization token is required"})
		c.Abort()
		return
	}
	// สามารถเพิ่ม logic การตรวจสอบ token ได้ที่นี่ เช่น ตรวจสอบ JWT หรือ token ในฐานข้อมูล
	c.Next()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// ตั้งค่า CORS ให้รองรับโดเมนที่ต้องการ
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// เพิ่ม authMiddleware สำหรับการตรวจสอบ token ก่อนใช้งาน routes ที่ต้องการ
	r.Use(authMiddleware)

	// กำหนด routes สำหรับ Brands
	brandRoutes := r.Group("/brands")
	{
		brandRoutes.GET("/", controllers.GetBrands)
		brandRoutes.GET("/:id", controllers.GetBrand)
		brandRoutes.POST("/", controllers.CreateBrand)
		brandRoutes.PUT("/:id", controllers.UpdateBrand)
		brandRoutes.DELETE("/:id", controllers.DeleteBrand)
	}

	// กำหนด routes สำหรับ Models
	modelRoutes := r.Group("/models")
	{
		modelRoutes.GET("/", controllers.GetModels)
		modelRoutes.GET("/:id", controllers.GetModel)
		modelRoutes.POST("/", controllers.CreateModel)
		modelRoutes.PUT("/:id", controllers.UpdateModel)
		modelRoutes.DELETE("/:id", controllers.DeleteModel)
	}

	// กำหนด routes สำหรับ Phones
	phoneRoutes := r.Group("/phones")
	{
		phoneRoutes.GET("/", controllers.GetPhones)
		phoneRoutes.GET("/:id", controllers.GetPhone)
		phoneRoutes.POST("/", controllers.CreatePhone)
		phoneRoutes.PUT("/:id", controllers.UpdatePhone)
		phoneRoutes.DELETE("/:id", controllers.DeletePhone)
	}

	// กำหนด routes สำหรับ Conditions
	conditionRoutes := r.Group("/conditions")
	{
		conditionRoutes.GET("/", controllers.GetConditions)
		conditionRoutes.GET("/:id", controllers.GetCondition)
		conditionRoutes.POST("/", controllers.CreateCondition)
		conditionRoutes.PUT("/:id", controllers.UpdateCondition)
		conditionRoutes.DELETE("/:id", controllers.DeleteCondition)
	}

	// กำหนด routes สำหรับ Pricing Adjustments
	pricingadjuctmentRoutes := r.Group("/pricingadjuctments")
	{
		pricingadjuctmentRoutes.GET("/", controllers.GetPricingAdjustments)
		pricingadjuctmentRoutes.GET("/:id", controllers.GetPricingAdjustment)
		pricingadjuctmentRoutes.POST("/", controllers.CreatePricingAdjustment)
		pricingadjuctmentRoutes.PUT("/:id", controllers.UpdatePricingAdjustment)
		pricingadjuctmentRoutes.DELETE("/:id", controllers.DeletePricingAdjustment)
	}

	return r
}
