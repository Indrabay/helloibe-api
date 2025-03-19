package repository

import "github.com/indrabay/helloibe-api/pkg/modules/auth/entity"

func (r *UserRepo) GetByUsername(username string) (res *entity.User, err error) {
	err = r.readDB.Where("username = ?", username).First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
