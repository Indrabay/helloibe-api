package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/internal/repository"
)

type ProductUsecase interface {
	UploadProductPrice(c *gin.Context, value []byte) error
	GetProductPrice(barcode string, storeID int64) (entity.ProductPrice, error)
	GetProductPrices(entity.GetProductsParams) ([]entity.ProductPrice, error)
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
