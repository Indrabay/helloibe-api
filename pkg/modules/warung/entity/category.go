package entity

import "time"

const CategoryTableName = "categories"

type Category struct {
	ID          int64     `json:"id" gorm:"column:id"`
	Name        string    `json:"name" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy   string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   string    `json:"updated_by" gorm:"column:updated_by"`
}

func (Category) TableName() string {
	return CategoryTableName
}

