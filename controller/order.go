package controller

import (
	"kart-server/config"
	"kart-server/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(context *gin.Context) {
	var orderReq models.OrderRequest
	if err := context.ShouldBindJSON(&orderReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if context.GetHeader("api_key") != config.AppConfig.APIKey {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	var originalTotal float64 = 0
	var productIDs []uint
	var quantities = make(map[uint]int)

	for _, item := range orderReq.Items {
		productIDs = append(productIDs, item.ProductID)
		quantities[item.ProductID] = item.Quantity
	}

	var products []models.Product
	if err := config.DB.Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	for _, product := range products {
		originalTotal += product.Price * float64(quantities[product.ID])
	}

	finalTotal := originalTotal
	var appliedDiscounts []string

	if orderReq.DiscountCodes != "" {
		discountCodes := strings.Split(orderReq.DiscountCodes, ",")
		trimmedCodes := make([]string, 0, len(discountCodes))

		for _, code := range discountCodes {
			trimmedCodes = append(trimmedCodes, strings.TrimSpace(code))
		}

		var discounts []models.Discount
		if err := config.DB.Where("discount_code IN ?", trimmedCodes).Find(&discounts).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch discount codes"})
			return
		}

		if len(discounts) != len(trimmedCodes) {
			foundCodes := make(map[string]bool)
			for _, discount := range discounts {
				foundCodes[discount.DiscountCode] = true
			}

			var invalidCodes []string
			for _, code := range trimmedCodes {
				if !foundCodes[code] {
					invalidCodes = append(invalidCodes, code)
				}
			}

			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discount code(s): " + strings.Join(invalidCodes, ", ")})
			return
		}

		now := time.Now()
		discountsToUpdate := make([]models.Discount, 0)

		for _, discount := range discounts {
			if now.Before(discount.ValidFrom) || now.After(discount.ValidTo) || discount.RemainingCount <= 0 {
				var reason string
				if now.Before(discount.ValidFrom) {
					reason = "not yet active"
				} else if now.After(discount.ValidTo) {
					reason = "expired"
				} else {
					reason = "no remaining uses"
				}

				context.JSON(http.StatusBadRequest, gin.H{"error": "Discount code " + discount.DiscountCode + " is " + reason})
				return
			}

			discountAmount := originalTotal * (discount.DiscountValue / 100)
			finalTotal -= discountAmount
			appliedDiscounts = append(appliedDiscounts, discount.DiscountCode)

			discount.RemainingCount--
			discountsToUpdate = append(discountsToUpdate, discount)
		}

		if len(discountsToUpdate) > 0 {
			tx := config.DB.Begin()
			for _, discount := range discountsToUpdate {
				if err := tx.Model(&models.Discount{}).Where("id = ?", discount.ID).Update("remaining_count", discount.RemainingCount).Error; err != nil {
					tx.Rollback()
					context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update discount codes"})
					return
				}
			}
			tx.Commit()
		}
	}

	order := models.Order{
		OriginalTotal: originalTotal,
		DiscountCode:  strings.Join(appliedDiscounts, ","),
		FinalTotal:    finalTotal,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	if len(orderReq.Items) > 0 {
		var orderItems []models.OrderItem
		for _, item := range orderReq.Items {
			orderItems = append(orderItems, models.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			})
		}

		if err := config.DB.Create(&orderItems).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order items"})
			return
		}
	}

	if err := config.DB.Preload("Items").Preload("Products").First(&order, order.ID).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load order details"})
		return
	}

	context.JSON(http.StatusOK, order)
}
