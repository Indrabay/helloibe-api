package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"github.com/indrabay/helloibe-api/utils/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUc) Register(user *entity.User) error {
	hashedPassword, err := uc.hashPassword(user.Password)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_Register"),
			zap.String("step", "hashPassword"),
		)
		return err
	}

	user.Password = hashedPassword

	err = uc.UserRepo.Insert(user)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_Register"),
			zap.String("step", "Insert"),
			zap.Any("user", user),
		)
		return err
	}

	userStore := entity.UserStore{
		UserID:  user.ID,
		StoreID: user.StoreID,
	}

	err = uc.UserRepo.InsertUserStore(userStore)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_Register"),
			zap.String("step", "InsertUserStore"),
			zap.Any("userStore", userStore),
		)
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
