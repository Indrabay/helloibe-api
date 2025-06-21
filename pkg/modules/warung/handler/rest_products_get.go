package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (h *ProductHandler) GetProducts(c *gin.Context) {
	params := h.assignQueryParam(c)
	response, err := h.ProductUsecase.GetProductPricesPaginated(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get products",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Products retrieved successfully",
		"data":    response,
	})
}

func (h *ProductHandler) assignQueryParam(c *gin.Context) entity.GetProductsParams {
	var (
		result entity.GetProductsParams
		store  any
		exist  bool
	)

	if store, exist = c.Get("store_id"); !exist {
		store = c.Param("store_id")
	}

	storeID, _ := strconv.ParseInt(store.(string), 10, 64)
	result.StoreID = int(storeID)

	queryParam := c.Request.URL.Query()

	// Product search parameters
	barcode := queryParam["barcode"]
	if len(barcode) > 0 {
		result.Barcode = barcode[0]
	}

	sku := queryParam["sku"]
	if len(sku) > 0 {
		result.SKU = sku[0]
	}

	name := queryParam["name"]
	if len(name) > 0 {
		result.Name = name[0]
	}

	// Pagination parameters
	limit := queryParam["limit"]
	if len(limit) > 0 {
		if limitVal, err := strconv.Atoi(limit[0]); err == nil && limitVal > 0 {
			result.Limit = limitVal
		} else {
			result.Limit = 10 // default limit
		}
	} else {
		result.Limit = 10 // default limit
	}

	// Validate and set maximum limit
	if result.Limit > 100 {
		result.Limit = 100 // maximum limit
	}

	page := queryParam["page"]
	if len(page) > 0 {
		if pageVal, err := strconv.Atoi(page[0]); err == nil && pageVal > 0 {
			result.Page = pageVal
		} else {
			result.Page = 1 // default page
		}
	} else {
		result.Page = 1 // default page
	}

	offset := queryParam["offset"]
	if len(offset) > 0 {
		if offsetVal, err := strconv.Atoi(offset[0]); err == nil && offsetVal >= 0 {
			result.Offset = offsetVal
		} else {
			result.Offset = 0 // default offset
		}
	} else {
		// Calculate offset from page if not provided
		result.Offset = (result.Page - 1) * result.Limit
	}

	return result
}
