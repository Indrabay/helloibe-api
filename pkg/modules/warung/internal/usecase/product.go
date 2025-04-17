package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/internal/repository"
)

type ProductUsecase interface {
	UploadCompleteProduct([]byte) error
}

type ProductUc struct {
	ProductRepo repository.ProductRepository
	PriceRepo   repository.PriceRepository
}

func NewProductUsecase(productRepo repository.ProductRepository, priceRepo repository.PriceRepository) *ProductUc {
	return &ProductUc{
		ProductRepo: productRepo,
		PriceRepo:   priceRepo,
	}
}
