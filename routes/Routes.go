package routes

import (
	"api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			allowedOrigins := []string{"http://localhost:3000", "http://localhost:8080"}
			for _, o := range allowedOrigins {
				if origin == o {
					return true
				}
			}
			return false
		},
	}))

	brandRoutes := r.Group("/brands")
	{
		brandRoutes.GET("/", controllers.GetBrands)
		brandRoutes.GET("/:id", controllers.GetBrand)
		brandRoutes.POST("/", controllers.CreateBrand)
		brandRoutes.PUT("/:id", controllers.UpdateBrand)
		brandRoutes.DELETE("/:id", controllers.DeleteBrand)
	}

	modelRoutes := r.Group("/models")
	{
		modelRoutes.GET("/", controllers.GetModels)
		modelRoutes.GET("/:id", controllers.GetModel)
		modelRoutes.POST("/", controllers.CreateModel)
		modelRoutes.PUT("/:id", controllers.UpdateModel)
		modelRoutes.DELETE("/:id", controllers.DeleteModel)
	}
	phoneRoutes := r.Group("/phones")
	{
		phoneRoutes.GET("/", controllers.GetPhones)
		phoneRoutes.GET("/:id", controllers.GetPhone)
		phoneRoutes.POST("/", controllers.CreatePhone)
		phoneRoutes.PUT("/:id", controllers.UpdatePhone)
		phoneRoutes.DELETE("/:id", controllers.DeletePhone)
	}
	conditionRoutes := r.Group("/conditions")
	{
		conditionRoutes.GET("/", controllers.GetConditions)
		conditionRoutes.GET("/:id", controllers.GetCondition)
		conditionRoutes.POST("/", controllers.CreateCondition)
		conditionRoutes.PUT("/:id", controllers.UpdateCondition)
		conditionRoutes.DELETE("/:id", controllers.DeleteCondition)
	}
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
