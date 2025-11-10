package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) SingleProduct(c *gin.Context) {
	sku := c.Param("sku")
	store, ok := c.Get("store_id")
	if !ok {
		store = c.Param("store_id")
	}
	storeID, _ := strconv.ParseInt(store.(string), 10, 64)

	productPrice, err := h.ProductUsecase.GetProductPrice(sku, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, productPrice)
}
