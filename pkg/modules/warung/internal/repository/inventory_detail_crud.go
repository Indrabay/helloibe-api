package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm/clause"
)

const InventoryDetailBatchSize = 50

// Insert creates a new inventory detail record
func (r *InventoryDetailRepo) Insert(detail *entity.InventoryDetail) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Create(detail).Error
}

// InsertBulk creates multiple inventory detail records
func (r *InventoryDetailRepo) InsertBulk(details []entity.InventoryDetail) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"quantity", "expiry_date", "location", "status", "updated_by"}),
	}).CreateInBatches(&details, InventoryDetailBatchSize).Error
}

// GetByID retrieves an inventory detail by its ID
func (r *InventoryDetailRepo) GetByID(id string) (detail entity.InventoryDetail, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("id = ?", id).First(&detail).Error
	return detail, err
}

// GetByInventoryID retrieves all inventory details for a specific inventory
func (r *InventoryDetailRepo) GetByInventoryID(inventoryID string) (details []entity.InventoryDetail, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("inventory_id = ?", inventoryID).Find(&details).Error
	return details, err
}

// Update updates an existing inventory detail record
func (r *InventoryDetailRepo) Update(detail *entity.InventoryDetail) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Save(detail).Error
}

// Delete deletes an inventory detail by ID
func (r *InventoryDetailRepo) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Where("id = ?", id).Delete(&entity.InventoryDetail{}).Error
}
