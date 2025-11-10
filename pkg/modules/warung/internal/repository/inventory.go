package repository

import (
	"sync"

	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Insert(inventory *entity.Inventory) error
	GetByID(id string) (entity.Inventory, error)
	GetByProductID(productID string) (entity.Inventory, error)
	GetByFilter(params entity.GetInventoriesParams) ([]entity.Inventory, error)
	GetCount(params entity.GetInventoriesParams) (int64, error)
	Update(inventory *entity.Inventory) error
	UpdateInventoryQuantity(inventoryID string) error
}

type InventoryRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
	mutex   sync.Mutex
}

func NewInventoryRepository(writeDB, readDB *gorm.DB) *InventoryRepo {
	return &InventoryRepo{
		writeDB: writeDB,
		readDB:  readDB,
		mutex:   sync.Mutex{},
	}
}
