package repository

import "github.com/indrabay/helloibe-api/pkg/modules/auth/entity"

func (r *UserRepo) InsertUserStore(userStore entity.UserStore) error {
	return r.writeDB.Save(&userStore).Error
}
