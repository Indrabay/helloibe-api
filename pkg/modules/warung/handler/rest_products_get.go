package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (h *ProductHandler) GetProducts(c *gin.Context) {
	params := h.assignQueryParam(c)
	productPrices, err := h.ProductUsecase.GetProductPrices(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, productPrices)
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

	return result
}
