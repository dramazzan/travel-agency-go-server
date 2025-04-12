package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server-go/internal/services"
	"strconv"
)

type BasketHandler struct {
	service services.BasketService
}

func NewBasketHandler(service services.BasketService) *BasketHandler {
	return &BasketHandler{service: service}
}

//func (h *BasketHandler) CreateBasket(c *gin.Context) {
//	userID := c.GetUint("userID")
//	basket, err := h.service.CreateBasket(uint(userID))
//	if err != nil {
//		log.Fatalf("error while creating basket: %v", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, basket)
//}

func (h *BasketHandler) GetBasket(c *gin.Context) {
	userID := c.GetUint("userID")
	basket, err := h.service.GetBasketByUserID(userID)
	if err != nil {
		log.Fatalf("error while getting basket: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Корзина не найдена"})
		return
	}
	c.JSON(http.StatusOK, basket)
}

func (h *BasketHandler) AddTourOnBasket(c *gin.Context) {
	userID := c.GetUint("userID")
	fmt.Println(userID)

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
