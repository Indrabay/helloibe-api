package handler

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (h *InventoryHandler) Create(c *gin.Context) {
	var inventory entity.Inventory
	if err := c.BindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.InventoryUsecase.Create(c, &inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Inventory created successfully",
		"data":    inventory,
	})
}

func (h *InventoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var inventory entity.Inventory
	if err := c.BindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.InventoryUsecase.Update(c, id, &inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inventory updated successfully",
	})
}

func (h *InventoryHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	inventory, err := h.InventoryUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inventory retrieved successfully",
		"data":    inventory,
	})
}

func (h *InventoryHandler) GetByFilter(c *gin.Context) {
	params := h.assignQueryParam(c)
	response, err := h.InventoryUsecase.GetByFilter(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inventories retrieved successfully",
		"data":    response,
	})
}

func (h *InventoryHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	source, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to open file"})
		return
	}
	defer source.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, source)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read file"})
		return
	}

	err = h.InventoryUsecase.Upload(c, buf.Bytes(), file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "file uploaded successfully"})
}

func (h *InventoryHandler) assignQueryParam(c *gin.Context) entity.GetInventoriesParams {
	var result entity.GetInventoriesParams

	store, ok := c.Get("store_id")
	if !ok {
		store = c.Param("store_id")
	}
	storeID, _ := strconv.ParseInt(store.(string), 10, 64)
	result.StoreID = int(storeID)

	queryParam := c.Request.URL.Query()

	productID := queryParam["product_id"]
	if len(productID) > 0 {
		result.ProductID = productID[0]
	}

	// Pagination parameters
	limit := queryParam["limit"]
	if len(limit) > 0 {
		if limitVal, err := strconv.Atoi(limit[0]); err == nil && limitVal > 0 {
			result.Limit = limitVal
		} else {
			result.Limit = 10
		}
	} else {
		result.Limit = 10
	}

	if result.Limit > 100 {
		result.Limit = 100
	}

	page := queryParam["page"]
	if len(page) > 0 {
		if pageVal, err := strconv.Atoi(page[0]); err == nil && pageVal > 0 {
			result.Page = pageVal
		} else {
			result.Page = 1
		}
	} else {
		result.Page = 1
	}

	offset := queryParam["offset"]
	if len(offset) > 0 {
		if offsetVal, err := strconv.Atoi(offset[0]); err == nil && offsetVal >= 0 {
			result.Offset = offsetVal
		} else {
			result.Offset = 0
		}
	} else {
		result.Offset = (result.Page - 1) * result.Limit
	}

	return result
}
