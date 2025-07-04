package entity

import "time"

const RoleTableName = "roles"

type Role struct {
	ID        int64     `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Level     int64     `json:"level" gorm:"column:level"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Role) TableName() string {
	return RoleTableName
}
