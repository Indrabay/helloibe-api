package repository

import (
	"sync"

	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm"
)

type InventoryDetailRepository interface {
	Insert(detail *entity.InventoryDetail) error
	InsertBulk(details []entity.InventoryDetail) error
	GetByID(id string) (entity.InventoryDetail, error)
	GetByInventoryID(inventoryID string) ([]entity.InventoryDetail, error)
	Update(detail *entity.InventoryDetail) error
	Delete(id string) error
}

type InventoryDetailRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
	mutex   sync.Mutex
}

func NewInventoryDetailRepository(writeDB, readDB *gorm.DB) *InventoryDetailRepo {
	return &InventoryDetailRepo{
		writeDB: writeDB,
		readDB:  readDB,
		mutex:   sync.Mutex{},
	}
}
