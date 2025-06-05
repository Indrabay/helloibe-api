package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUc) Register(user *entity.User) error {
	hashedPassword, err := uc.hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = uc.UserRepo.Insert(user)
	if err != nil {
		return err
	}

	userStore := entity.UserStore{
		UserID:  user.ID,
		StoreID: user.StoreID,
	}

	err = uc.UserRepo.InsertUserStore(userStore)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUc) hashPassword(inputPassword string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 14)
	if err != nil {
		return "", err
	}

	return string(hashPassword), err
}
