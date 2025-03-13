package handlers

import (
	"7-solutions-challenges/internal/meat_count/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MeatCountHandler provides HTTP handlers for meat count endpoints
type MeatCountHandler struct {
	usecase *usecases.MeatCountUsecase
}

func NewMeatCountHandler(usecase *usecases.MeatCountUsecase) *MeatCountHandler {
	return &MeatCountHandler{usecase: usecase}
}

func (h *MeatCountHandler) GetBeefSummary(c *gin.Context) {
	meatCounts, err := h.usecase.GetMeatCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch meat count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"beef": meatCounts,
	})
}
