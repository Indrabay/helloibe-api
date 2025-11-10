package entity

import "time"

const InventoryTableName = "inventories"

type Inventory struct {
	ID        string    `json:"id" gorm:"column:id"`
	ProductID string    `json:"product_id" gorm:"column:product_id"`
	Quantity  float64   `json:"quantity" gorm:"column:quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy string    `json:"updated_by" gorm:"column:updated_by"`
}

func (Inventory) TableName() string {
	return InventoryTableName
}
