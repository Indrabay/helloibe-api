package entity

import "time"

const StoreTableName = "stores"

type Store struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Store) TableName() string {
	return StoreTableName
}
