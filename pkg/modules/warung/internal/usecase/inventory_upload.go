package usecase

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"github.com/xuri/excelize/v2"
)

func (uc *InventoryUc) Upload(ctx *gin.Context, value []byte, filename string) error {
	username, exist := ctx.Get("username")
	if !exist {
		return fmt.Errorf("you need to login to use this endpoint")
	}

	_, exist = ctx.Get("store_id")
	if !exist {
		return fmt.Errorf("store_id is required")
	}

	// Detect file type by extension
	ext := strings.ToLower(filepath.Ext(filename))
	var records [][]string
	var err error

	if ext == ".xlsx" || ext == ".xls" {
		records, err = uc.parseXLSX(value)
		if err != nil {
			return fmt.Errorf("failed to parse XLSX file: %v", err)
		}
	} else if ext == ".csv" {
		records, err = uc.parseCSV(value)
		if err != nil {
			return fmt.Errorf("failed to parse CSV file: %v", err)
		}
	} else {
		return fmt.Errorf("unsupported file format. Only CSV and XLSX files are supported")
	}

	if len(records) < 2 {
		return fmt.Errorf("file you are uploading is empty or has no data rows")
	}

	// Expected format: product_id, quantity, expiry_date (optional), location (optional), status (optional)
	// Skip header (first row)
	for _, record := range records[1:] {
		if len(record) < 2 {
			continue
		}

		productID := record[0]
		quantity, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return fmt.Errorf("invalid quantity: %s", record[1])
		}

		// Check if inventory exists for this product
		inventory, err := uc.InventoryRepo.GetByProductID(productID)
		if err != nil {
			// Create new inventory if doesn't exist
			inventory = entity.Inventory{
				ID:        uuid.NewString(),
				ProductID: productID,
				Quantity:  0,
				CreatedBy: username.(string),
				UpdatedBy: username.(string),
			}
			err = uc.InventoryRepo.Insert(&inventory)
			if err != nil {
				return fmt.Errorf("failed to create inventory: %v", err)
			}
		}

		// Parse optional fields
		var expiryDate *time.Time
		if len(record) > 2 && record[2] != "" {
			parsedDate, err := time.Parse("2006-01-02", record[2])
			if err != nil {
				// Try other date formats
				parsedDate, err = time.Parse("2006/01/02", record[2])
				if err != nil {
					return fmt.Errorf("invalid expiry_date format: %s", record[2])
				}
			}
			expiryDate = &parsedDate
		}

		location := ""
		if len(record) > 3 {
			location = record[3]
		}

		status := "active"
		if len(record) > 4 && record[4] != "" {
			status = record[4]
		}

		// Create inventory detail
		detail := entity.InventoryDetail{
			ID:          uuid.NewString(),
			InventoryID: inventory.ID,
			Quantity:    quantity,
			ExpiryDate:  expiryDate,
			Location:    location,
			Status:      status,
			CreatedBy:   username.(string),
			UpdatedBy:   username.(string),
		}

		err = uc.InventoryDetailRepo.Insert(&detail)
		if err != nil {
			return fmt.Errorf("failed to create inventory detail: %v", err)
		}

		// Update inventory quantity
		err = uc.InventoryRepo.UpdateInventoryQuantity(inventory.ID)
		if err != nil {
			return fmt.Errorf("failed to update inventory quantity: %v", err)
		}
	}

	return nil
}

// parseCSV parses CSV file content and returns records
func (uc *InventoryUc) parseCSV(value []byte) ([][]string, error) {
	reader := csv.NewReader(bytes.NewReader(value))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

// parseXLSX parses XLSX file content and returns records
func (uc *InventoryUc) parseXLSX(value []byte) ([][]string, error) {
	f, err := excelize.OpenReader(bytes.NewReader(value))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, fmt.Errorf("no sheets found in Excel file")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	records := make([][]string, len(rows))
	for i, row := range rows {
		records[i] = row
	}

	return records, nil
}
