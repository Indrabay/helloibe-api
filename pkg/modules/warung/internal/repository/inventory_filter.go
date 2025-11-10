package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

// GetByFilter retrieves inventories with filtering and pagination
func (r *InventoryRepo) GetByFilter(params entity.GetInventoriesParams) (inventories []entity.Inventory, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	query := r.readDB.Model(&entity.Inventory{})

	// Join with products to filter by store_id
	if params.StoreID > 0 {
		query = query.Joins("JOIN products ON inventories.product_id = products.id").
			Where("products.store_id = ?", params.StoreID)
	}

	if params.ProductID != "" {
		query = query.Where("inventories.product_id = ?", params.ProductID)
	}

	// Add ORDER BY for consistent pagination
	query = query.Order("inventories.created_at DESC")

	// Add pagination
	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	if params.Offset > 0 {
		query = query.Offset(params.Offset)
	}

	err = query.Find(&inventories).Error
	return inventories, err
}

// GetCount returns the total count of inventories matching the search criteria
func (r *InventoryRepo) GetCount(params entity.GetInventoriesParams) (count int64, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	query := r.readDB.Model(&entity.Inventory{})

	// Join with products to filter by store_id
	if params.StoreID > 0 {
		query = query.Joins("JOIN products ON inventories.product_id = products.id").
			Where("products.store_id = ?", params.StoreID)
	}

	if params.ProductID != "" {
		query = query.Where("inventories.product_id = ?", params.ProductID)
	}

	err = query.Count(&count).Error
	return count, err
}
