package repository

import (
	"sync"

	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm"
)

type PriceRepository interface {
	InsertBulk(prices []entity.Price) error
	GetPrice(productID string) (entity.Price, error)
	GetPrices(productIDs []string) ([]entity.Price, error)
}

type PriceRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
	mutex   sync.Mutex
}

func NewPriceRepository(writeDB, readDB *gorm.DB) *PriceRepo {
	return &PriceRepo{
		writeDB: writeDB,
		readDB:  readDB,
		mutex:   sync.Mutex{},
	}
}
