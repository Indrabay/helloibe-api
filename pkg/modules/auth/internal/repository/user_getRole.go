package repository

import "github.com/indrabay/helloibe-api/pkg/modules/auth/entity"

func (r *UserRepo) GetRole(roleID int) (res *entity.Role, err error) {
	err = r.readDB.Where("id = ?", roleID).First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
