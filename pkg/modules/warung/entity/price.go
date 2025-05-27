package entity

import "time"

const PriceTableName = "prices"

type Price struct {
	ID            string    `json:"id" gorm:"column:id"`
	ProductID     string    `json:"product_id" gorm:"column:product_id"`
	SellingPrice  float64   `json:"selling_price" gorm:"column:selling_price"`
	PurchasePrice float64   `json:"purchase_price" gorm:"column:purchase_price"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy     string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy     string    `json:"updated_by" gorm:"column:updated_by"`
}

func (Price) TableName() string {
	return PriceTableName
}
