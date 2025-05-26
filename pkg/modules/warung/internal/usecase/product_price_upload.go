package usecase

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (u *ProductUc) UploadProductPrice(ctx *gin.Context, value []byte) error {
	var (
		products    []entity.Product
		priceLookup = make(map[string]entity.Price)
		prices      []entity.Price
		barcodes    []string
	)

	username, exist := ctx.Get("username")
	if !exist {
		return fmt.Errorf("you need to login to use this endpoint")
	}
	reader := csv.NewReader(bytes.NewReader(value))

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records[1:]) < 1 {
		return fmt.Errorf("file you are uploading is empty")
	}

	// skip header
	for _, record := range records[1:] {
		product := entity.Product{
			ID:        uuid.NewString(),
			Barcode:   record[1],
			Name:      record[0],
			SKU:       record[2],
			CreatedBy: username.(string),
			UpdatedBy: username.(string),
		}

		barcodes = append(barcodes, record[1])

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
		}

		priceLookup[product.Barcode] = price
	}

	err = u.ProductRepo.InsertBulk(products)
	if err != nil {
		return err
	}

	lastUpdatedProducts, err := u.ProductRepo.GetProductsByBarcode(barcodes)
	if err != nil {
		return err
	}

	for _, product := range lastUpdatedProducts {
		if _, ok := priceLookup[product.Barcode]; !ok {
			continue
		}

		price := priceLookup[product.Barcode]
		price.ProductID = product.ID

		prices = append(prices, price)
	}

	err = u.PriceRepo.InsertBulk(prices)
	if err != nil {
		return err
	}

	return nil
}
