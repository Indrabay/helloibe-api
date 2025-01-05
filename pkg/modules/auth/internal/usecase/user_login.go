package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"github.com/indrabay/helloibe-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUc) Login(username, password string) (jwt string, err error) {
	var (
		user  entity.User
		token string
	)

	user, err = uc.UserRepo.GetByUsername(username)
	if err != nil {
		uc.Logger.Error("usecase", "UserUC_Login-GetUser", err)
		return "", utils.ErrUserNotFound
	}

	err = uc.validatePassword(user.Password, password)
	if err != nil {
		uc.Logger.Error("usecase", "UserUC_Login-validatePassword", err)
		return "", utils.ErrUserNotFound
	}

	claims := utils.JWTClaim{
		Username: user.Username,
		Name:     user.Name,
		Role:     user.Role,
	}

	token, err = uc.JWTUtils.CreateToken(claims)
	if err != nil {
		uc.Logger.Error("usecase", "UserUC_Login-createToken", err)
		return "", err
	}

	return token, nil
}

func (uc *UserUc) validatePassword(dbPassword, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword))
	if err != nil {
		return err
	}

	return nil
}
