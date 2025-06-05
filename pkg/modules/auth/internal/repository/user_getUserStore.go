package repository

import "github.com/indrabay/helloibe-api/pkg/modules/auth/entity"

func (r *UserRepo) GetUserStores(userID int64) (res []entity.UserStore, err error) {
	err = r.readDB.Where("user_id = ?", userID).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
