package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertBulk(products []entity.Product) error
	GetProductsByBarcode(barcodes []string) ([]entity.Product, error)
	GetProductByBarcodeStore(barcode string, storeID int64) (entity.Product, error)
	GetProducts(entity.GetProductsParams) ([]entity.Product, error)
}

type ProductRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewProductRepository(writeDB, readDB *gorm.DB) *ProductRepo {
	return &ProductRepo{
		writeDB: writeDB,
		readDB:  readDB,
	}
}
