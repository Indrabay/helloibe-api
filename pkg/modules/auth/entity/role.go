package entity

const RoleTableName = "role"

type Role struct {
	ID        int64  `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

func (Role) TableName() string {
	return RoleTableName
}
