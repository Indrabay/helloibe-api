package usecase

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"github.com/xuri/excelize/v2"
)

func (u *ProductUc) UploadProductPrice(ctx *gin.Context, value []byte, filename string) error {
	var (
		products    []entity.Product
		priceLookup = make(map[string]entity.Price)
		prices      []entity.Price
		skus        []string
	)

	username, exist := ctx.Get("username")
	if !exist {
		return fmt.Errorf("you need to login to use this endpoint")
	}

	storeID, exist := ctx.Get("store_id")
	if !exist {
		return fmt.Errorf("store_id is required")
	}

	storeIDInt, ok := storeID.(int64)
	if !ok {
		// Try to convert from string if needed
		if storeIDStr, ok := storeID.(string); ok {
			var err error
			storeIDInt, err = strconv.ParseInt(storeIDStr, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid store_id format")
			}
		} else {
			return fmt.Errorf("invalid store_id type")
		}
	}

	// Detect file type by extension
	ext := strings.ToLower(filepath.Ext(filename))
	var records [][]string
	var err error

	if ext == ".xlsx" || ext == ".xls" {
		// Parse XLSX file
		records, err = u.parseXLSX(value)
		if err != nil {
			return fmt.Errorf("failed to parse XLSX file: %v", err)
		}
	} else if ext == ".csv" {
		// Parse CSV file
		records, err = u.parseCSV(value)
		if err != nil {
			return fmt.Errorf("failed to parse CSV file: %v", err)
		}
	} else {
		return fmt.Errorf("unsupported file format. Only CSV and XLSX files are supported")
	}

	if len(records) < 2 {
		return fmt.Errorf("file you are uploading is empty or has no data rows")
	}

	// skip header (first row)
	// Expected format: name, category_id, sku (optional), purchase_price, selling_price
	// Track category counts for SKU generation
	categoryCounts := make(map[int64]int64)

	for _, record := range records[1:] {
		categoryID, err := strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid category_id: %s", record[1])
		}

		// Get SKU from CSV or generate if empty
		sku := record[2]
		if sku == "" {
			// Generate SKU based on category_id
			_, exists := categoryCounts[categoryID]
			if !exists {
				// Get current count from database for this category and store
				dbCount, err := u.ProductRepo.GetProductCountByCategoryAndStore(categoryID, storeIDInt)
				if err != nil {
					return fmt.Errorf("failed to get product count for category %d: %v", categoryID, err)
				}
				categoryCounts[categoryID] = dbCount
			}
			// Increment count for this category
			categoryCounts[categoryID]++
			// Generate SKU: CAT{category_id}-{sequence}
			sku = fmt.Sprintf("CAT%d-%04d", categoryID, categoryCounts[categoryID])
		}

		product := entity.Product{
			ID:         uuid.NewString(),
			Name:       record[0],
			CategoryID: categoryID,
			SKU:        sku,
			StoreID:    int(storeIDInt),
			CreatedBy:  username.(string),
			UpdatedBy:  username.(string),
		}

		skus = append(skus, sku)

		products = append(products, product)
		purchasePrice, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return err
		}
		sellingPrice, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return err
		}

		price := entity.Price{
			ID:            uuid.NewString(),
			PurchasePrice: purchasePrice,
			SellingPrice:  sellingPrice,
			UpdatedBy:     username.(string),
			CreatedBy:     username.(string),
		}

		priceLookup[product.SKU] = price
	}

	err = u.ProductRepo.InsertBulk(products)
	if err != nil {
		return err
	}

	// Link prices to products using the product IDs we just created
	for i, product := range products {
		if price, ok := priceLookup[product.SKU]; ok {
			price.ProductID = product.ID
			prices = append(prices, price)
		} else if i < len(skus) {
			// Fallback: try to find price by SKU from the list
			sku := skus[i]
			if price, ok := priceLookup[sku]; ok {
				price.ProductID = product.ID
				prices = append(prices, price)
			}
		}
	}

	err = u.PriceRepo.InsertBulk(prices)
	if err != nil {
		return err
	}

	return nil
}

// parseCSV parses CSV file content and returns records
func (u *ProductUc) parseCSV(value []byte) ([][]string, error) {
	reader := csv.NewReader(bytes.NewReader(value))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

// parseXLSX parses XLSX file content and returns records
func (u *ProductUc) parseXLSX(value []byte) ([][]string, error) {
	f, err := excelize.OpenReader(bytes.NewReader(value))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Get the first sheet name
	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, fmt.Errorf("no sheets found in Excel file")
	}

	// Get all rows from the first sheet
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	// Convert to [][]string format (same as CSV)
	records := make([][]string, len(rows))
	for i, row := range rows {
		records[i] = row
	}

	return records, nil
}
