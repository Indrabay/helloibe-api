package repository

import "github.com/indrabay/helloibe-api/pkg/modules/auth/entity"

// GetUserStoreCount returns the count of stores associated with a user
func (r *UserRepo) GetUserStoreCount(userID int64) (count int64, err error) {
	err = r.readDB.Model(&entity.UserStore{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

