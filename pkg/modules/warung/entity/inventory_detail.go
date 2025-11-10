package entity

import "time"

const InventoryDetailTableName = "inventory_details"

type InventoryDetail struct {
	ID          string     `json:"id" gorm:"column:id"`
	InventoryID string     `json:"inventory_id" gorm:"column:inventory_id"`
	Quantity    float64    `json:"quantity" gorm:"column:quantity"`
	ExpiryDate  *time.Time `json:"expiry_date" gorm:"column:expiry_date"`
	Location    string     `json:"location" gorm:"column:location"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy   string     `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   string     `json:"updated_by" gorm:"column:updated_by"`
}

func (InventoryDetail) TableName() string {
	return InventoryDetailTableName
}
