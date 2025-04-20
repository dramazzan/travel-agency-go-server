package handlers

import (
	"log"
	"net/http"
	"server-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BasketHandler struct {
	service services.BasketService
}

func NewBasketHandler(service services.BasketService) *BasketHandler {
	return &BasketHandler{service: service}
}

func (h *BasketHandler) GetBasket(c *gin.Context) {
	userID := c.GetUint("userID")
	basket, err := h.service.GetBasketByUserID(userID)
	if err != nil {
		log.Printf("Error while getting basket: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Корзина не найдена"})
		return
	}
	c.JSON(http.StatusOK, basket)
}

func (h *BasketHandler) AddTourToBasket(c *gin.Context) {
	userID := c.GetUint("userID")
	tourID, err := strconv.Atoi(c.Param("tourID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	basket, err := h.service.GetBasketByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get or create basket"})
		return
	}

	err = h.service.AddTourToBasket(basket.ID, uint(tourID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour added to basket successfully"})
}

func (h *BasketHandler) RemoveTourFromBasket(c *gin.Context) {
	userID := c.GetUint("userID")
	tourID, err := strconv.Atoi(c.Param("tourID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	basket, err := h.service.GetBasketByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get basket"})
		return
	}

	err = h.service.RemoveTourFromBasket(basket.ID, uint(tourID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tour removed from basket successfully"})
}
