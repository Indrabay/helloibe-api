package repository

import "github.com/indrabay/helloibe-api/pkg/modules/auth/entity"

func (r *UserRepo) Insert(user *entity.User) error {
	err := r.writeDB.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}
