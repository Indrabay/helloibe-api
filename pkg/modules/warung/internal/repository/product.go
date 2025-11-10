package repository

import (
	"sync"

	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertBulk(products []entity.Product) error
	GetProductBySKU(sku string, storeID int64) (entity.Product, error)
	GetProducts(entity.GetProductsParams) ([]entity.Product, error)
	GetStoreStatistics(storeID int64) (entity.StoreStatistics, error)
	GetProductCountByStore(storeID int64) (int64, error)
	GetPriceCountByStore(storeID int64) (int64, error)
	GetProductsCount(entity.GetProductsParams) (int64, error)
	GetProductCountByCategoryAndStore(categoryID int64, storeID int64) (int64, error)
}

type ProductRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
	mutex   sync.Mutex
}

func NewProductRepository(writeDB, readDB *gorm.DB) *ProductRepo {
	return &ProductRepo{
		writeDB: writeDB,
		readDB:  readDB,
		mutex:   sync.Mutex{},
	}
}
