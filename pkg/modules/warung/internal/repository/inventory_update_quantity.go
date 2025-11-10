package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

// UpdateInventoryQuantity calculates and updates inventory.quantity as the sum of all inventory_detail.quantity
// for the given inventory_id. Only counts inventory_details with status = 'active'
func (r *InventoryRepo) UpdateInventoryQuantity(inventoryID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var totalQuantity float64
	err := r.readDB.Model(&entity.InventoryDetail{}).
		Where("inventory_id = ? AND status = ?", inventoryID, "active").
		Select("COALESCE(SUM(quantity), 0)").
		Scan(&totalQuantity).Error
	if err != nil {
		return err
	}

	err = r.writeDB.Model(&entity.Inventory{}).
		Where("id = ?", inventoryID).
		Update("quantity", totalQuantity).Error

	return err
}
