package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) GetStoreStatistics(c *gin.Context) {
	store, ok := c.Get("store_id")
	if !ok {
		store = c.Param("store_id")
	}
	storeID, _ := strconv.ParseInt(store.(string), 10, 64)

	statistics, err := h.ProductUsecase.GetStoreStatistics(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get store statistics",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Store statistics retrieved successfully",
		"data":    statistics,
	})
}
