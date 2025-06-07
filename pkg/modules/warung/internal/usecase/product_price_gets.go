package usecase

import (
	"errors"

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
