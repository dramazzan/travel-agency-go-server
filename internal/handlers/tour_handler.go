package handlers

import (
	"net/http"
	"server-go/internal/models"
	"server-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TourHandler struct {
	service services.TourService
}

type tourRequest struct {
	Category uint `json:"category" binding:"required"`
}

func NewTourHandler(service services.TourService) *TourHandler {
	return &TourHandler{service: service}
}

func (h *TourHandler) GetAllTours(c *gin.Context) {
	tours, err := h.service.GetAllTours()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tours)
}

func (h *TourHandler) GetTourByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	tour, err := h.service.GetTourByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тур не найден"})
		return
	}

	c.JSON(http.StatusOK, tour)
}

func (h *TourHandler) CreateTour(c *gin.Context) {
	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tour.ID = 0

	if err := h.service.CreateTour(&tour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tour)
}

func (h *TourHandler) UpdateTour(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tour.ID = uint(id)
	if err := h.service.UpdateTour(&tour); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tour)
}

func (h *TourHandler) DeleteTour(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	if err := h.service.DeleteTour(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Тур успешно удален"})
}

func (h *TourHandler) GetTourByCategory(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 'category' обязателен"})
		return
	}

	tours, err := h.service.GetTourByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tours)
}
