package controller

import (
	"fmt"
	"kart-server/config"
	"kart-server/models"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListProducts(context *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	context.JSON(http.StatusOK, products)
}

func GetProduct(context *gin.Context) {
	id := context.Param("productId")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	context.JSON(http.StatusOK, product)
}

func CreateProduct(context *gin.Context) {
	if context.GetHeader("api_key") != config.AppConfig.APIKey {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	name := context.PostForm("name")
	price := context.PostForm("price")
	category := context.PostForm("category")

	if name == "" || price == "" || category == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	file, err := context.FormFile("image")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFilename := uuid.New().String() + extension
	imagePath := filepath.Join("public/images", newFilename)

	if err := context.SaveUploadedFile(file, imagePath); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	imageURL := fmt.Sprintf("/images/%s", newFilename)

	product := models.Product{
		Name:      name,
		Price:     parseFloat(price),
		Category:  category,
		ImageURL:  imageURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := config.DB.Create(&product).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	context.JSON(http.StatusCreated, product)
}

func parseFloat(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}
