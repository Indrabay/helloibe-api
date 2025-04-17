package entity

const UserStoreTableName = "user_store"

type UserStore struct {
	UserID  int64 `json:"user_id" gorm:"column:user_id"`
	StoreID int   `json:"store_id" gorm:"column:store_id"`
}

func (UserStore) TableName() string {
	return UserStoreTableName
}
