package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

// Insert creates a new inventory record
func (r *InventoryRepo) Insert(inventory *entity.Inventory) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Create(inventory).Error
}

// GetByID retrieves an inventory by its ID
func (r *InventoryRepo) GetByID(id string) (inventory entity.Inventory, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("id = ?", id).First(&inventory).Error
	return inventory, err
}

// GetByProductID retrieves an inventory by product_id
func (r *InventoryRepo) GetByProductID(productID string) (inventory entity.Inventory, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("product_id = ?", productID).First(&inventory).Error
	return inventory, err
}

// Update updates an existing inventory record
func (r *InventoryRepo) Update(inventory *entity.Inventory) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Save(inventory).Error
}

