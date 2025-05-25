package controller

import (
	"kart-server/config"
	"kart-server/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDiscount(context *gin.Context) {
	code := context.Param("discountCode")

	var discount models.Discount
	if err := config.DB.Where("discount_code = ?", code).First(&discount).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Discount code not found"})
		return
	}

	now := time.Now()
	if now.Before(discount.ValidFrom) || now.After(discount.ValidTo) || discount.RemainingCount <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Discount code is expired or invalid"})
		return
	}

	response := models.DiscountResponse{
		DiscountCode:   discount.DiscountCode,
		DiscountValue:  discount.DiscountValue,
		ValidFrom:      discount.ValidFrom,
		ValidTo:        discount.ValidTo,
		RemainingCount: discount.RemainingCount,
	}

	context.JSON(http.StatusOK, response)
}
