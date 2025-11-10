package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func (h *ProductHandler) GetAllProductPrices(c *gin.Context) {
	store, ok := c.Get("store_id")
	if !ok {
		store = c.Param("store_id")
	}
	storeID, err := strconv.ParseInt(store.(string), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid store_id",
		})
		return
	}

	productPrices, err := h.ProductUsecase.GetAllProductPricesByStore(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get product prices",
			"error":   err.Error(),
		})
		return
	}

	// Create a new Excel file
	f := excelize.NewFile()
	defer f.Close()

	sheetName := "Product Prices"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create Excel sheet",
		})
		return
	}

	// Delete default sheet
	f.DeleteSheet("Sheet1")

	// Set active sheet
	f.SetActiveSheet(index)

	// Set headers
	headers := []string{"ID", "Name", "SKU", "Category ID", "Store ID", "Purchase Price", "Selling Price", "Created At", "Updated At"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Style headers
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#E0E0E0"}, Pattern: 1},
	})
	if err == nil {
		f.SetCellStyle(sheetName, "A1", fmt.Sprintf("%c1", 'A'+len(headers)-1), headerStyle)
	}

	// Add data rows
	for i, productPrice := range productPrices {
		row := i + 2 // Start from row 2 (row 1 is headers)
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), productPrice.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), productPrice.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), productPrice.SKU)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), productPrice.CategoryID)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), productPrice.StoreID)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), productPrice.PurchasePrice)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), productPrice.SellingPrice)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), productPrice.CreatedAt.Format("2006-01-02 15:04:05"))
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), productPrice.UpdatedAt.Format("2006-01-02 15:04:05"))
	}

	// Set column widths
	colWidths := map[string]float64{
		"A": 30, // ID
		"B": 30, // Name
		"C": 20, // SKU
		"D": 15, // Category ID
		"E": 12, // Store ID
		"F": 18, // Purchase Price
		"G": 18, // Selling Price
		"H": 20, // Created At
		"I": 20, // Updated At
	}
	for col, width := range colWidths {
		f.SetColWidth(sheetName, col, col, width)
	}

	// Set response headers
	filename := fmt.Sprintf("product_prices_store_%d_%s.xlsx", storeID, time.Now().Format("20060102_150405"))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Write Excel file to response
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate Excel file",
		})
		return
	}
}
