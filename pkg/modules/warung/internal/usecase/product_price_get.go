package usecase

import (
	"errors"

	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (uc *ProductUc) GetProductPrice(barcode string, storeID int64) (res entity.ProductPrice, err error) {
	var (
		product entity.Product
		price   entity.Price
	)
	product, err = uc.ProductRepo.GetProductByBarcodeStore(barcode, storeID)
	if err != nil {
		return res, err
	}

	if product == (entity.Product{}) {
		return res, errors.New("product not found")
	}

	price, err = uc.PriceRepo.GetPrice(product.ID)
	if err != nil {
		return res, err
	}

	res.Product = product
	res.PurchasePrice = price.PurchasePrice
	res.SellingPrice = price.SellingPrice

	return res, nil
}
