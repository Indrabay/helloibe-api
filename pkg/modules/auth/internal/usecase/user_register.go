package usecase

import (
	"time"

	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"github.com/indrabay/helloibe-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUc) Register(user *entity.User) error {
	hashedPassword, err := uc.hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.CreatedAt = time.Now().Format(utils.DateFormat)
	user.UpdatedAt = time.Now().Format(utils.DateFormat)

	err = uc.UserRepo.Insert(user)
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
