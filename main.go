package main

import (
	"kart-server/config"
	"kart-server/controller"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     config.AppConfig.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "api_key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.Static("/images", "./public/images")

	api := engine.Group("/api")

	api.GET("/product", controller.ListProducts)
	api.GET("/product/:productId", controller.GetProduct)
	api.POST("/product", controller.CreateProduct)

	api.GET("/discount/:discountCode", controller.GetDiscount)

	api.POST("/order", controller.PlaceOrder)

	log.Printf("Starting server on port %s", config.AppConfig.AppPort)
	engine.Run(":" + config.AppConfig.AppPort)
}
