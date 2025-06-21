package usecase

import (
	"errors"
	"math"

	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (uc *ProductUc) GetProductPrices(params entity.GetProductsParams) (res []entity.ProductPrice, err error) {
	var (
		products           []entity.Product
		prices             []entity.Price
		productIDs         []string
		productPriceLookup = make(map[string]entity.ProductPrice)
	)
	products, err = uc.ProductRepo.GetProducts(params)
	if err != nil {
		return res, err
	}

	if len(products) == 0 {
		return res, errors.New("record not found - product")
	}

	for _, product := range products {
		productIDs = append(productIDs, product.ID)
		productPrice := entity.ProductPrice{
			Product: product,
		}
		productPriceLookup[product.ID] = productPrice
	}

	prices, err = uc.PriceRepo.GetPrices(productIDs)
	if err != nil {
		return res, err
	}

	if len(prices) == 0 {
		return res, errors.New("record not found - price")
	}

	for _, price := range prices {
		if productPrice, ok := productPriceLookup[price.ProductID]; ok {
			productPriceObj := productPrice
			productPriceObj.PurchasePrice = price.PurchasePrice
			productPriceObj.SellingPrice = price.SellingPrice

			productPriceLookup[price.ProductID] = productPriceObj
		}
	}

	for _, v := range productPriceLookup {
		res = append(res, v)
	}

	return res, nil
}

// GetProductPricesPaginated returns paginated product prices with metadata
func (uc *ProductUc) GetProductPricesPaginated(params entity.GetProductsParams) (response entity.PaginatedProductResponse, err error) {
	// Get total count for pagination metadata
	totalItems, err := uc.ProductRepo.GetProductsCount(params)
	if err != nil {
		return response, err
	}

	// Get paginated products
	productPrices, err := uc.GetProductPrices(params)
	if err != nil {
		return response, err
	}

	// Calculate pagination metadata
	totalPages := int(math.Ceil(float64(totalItems) / float64(params.Limit)))
	if totalPages == 0 {
		totalPages = 1
	}

	hasNext := params.Page < totalPages
	hasPrev := params.Page > 1

	response = entity.PaginatedProductResponse{
		Data: productPrices,
		Pagination: entity.PaginationMeta{
			CurrentPage:  params.Page,
			TotalPages:   totalPages,
			TotalItems:   totalItems,
			ItemsPerPage: params.Limit,
			HasNext:      hasNext,
			HasPrev:      hasPrev,
		},
	}

	return response, nil
}
