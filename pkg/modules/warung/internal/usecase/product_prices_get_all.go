package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

// GetAllProductPricesByStore retrieves all product prices for a specific store
func (uc *ProductUc) GetAllProductPricesByStore(storeID int64) (res []entity.ProductPrice, err error) {
	var (
		products           []entity.Product
		prices             []entity.Price
		productIDs         []string
		productPriceLookup = make(map[string]entity.ProductPrice)
	)

	// Get all products for the store
	params := entity.GetProductsParams{
		StoreID: int(storeID),
		Limit:   10000, // Large limit to get all products
		Page:    1,
		Offset:  0,
	}

	products, err = uc.ProductRepo.GetProducts(params)
	if err != nil {
		return res, err
	}

	if len(products) == 0 {
		return res, nil // Return empty array instead of error if no products
	}

	// Build product price lookup map
	for _, product := range products {
		productIDs = append(productIDs, product.ID)
		productPrice := entity.ProductPrice{
			Product: product,
		}
		productPriceLookup[product.ID] = productPrice
	}

	// Get all prices for these products
	prices, err = uc.PriceRepo.GetPrices(productIDs)
	if err != nil {
		return res, err
	}

	// Combine products with prices
	for _, price := range prices {
		if productPrice, ok := productPriceLookup[price.ProductID]; ok {
			productPriceObj := productPrice
			productPriceObj.PurchasePrice = price.PurchasePrice
			productPriceObj.SellingPrice = price.SellingPrice
			productPriceLookup[price.ProductID] = productPriceObj
		}
	}

	// Convert map to slice
	for _, v := range productPriceLookup {
		res = append(res, v)
	}

	return res, nil
}

