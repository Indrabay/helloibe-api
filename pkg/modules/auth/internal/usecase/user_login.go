package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"github.com/indrabay/helloibe-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUc) Login(username, password string) (*entity.CompleteUser, string, error) {
	var (
		userDetail *entity.User
		err        error
		token      string
	)

	userDetail, err = uc.UserRepo.GetByUsername(username)
	if err != nil {
		uc.Logger.Error("usecase", "UserUC_Login-GetUser", err)
		return nil, "", utils.ErrUserNotFound
	}

	err = uc.validatePassword(userDetail.Password, password)
	if err != nil {
		uc.Logger.Error("usecase", "UserUC_Login-validatePassword", err)
		return nil, "", utils.ErrUserNotFound
	}

	claims := utils.JWTClaim{
		Username: userDetail.Username,
		Name:     userDetail.Name,
		Role:     userDetail.Role,
	}

	token, err = uc.JWTUtils.CreateToken(claims)
	if err != nil {
		uc.Logger.Error("usecase", "UserUC_Login-createToken", err)
		return nil, "", err
	}

	role, err := uc.UserRepo.GetRole(userDetail.Role)
	if err != nil {
		uc.Logger.Error("usecase", "UserUC_Login-getRole", err)
		return nil, "", err
	}

	userStores, err := uc.UserRepo.GetUserStores(userDetail.ID)
	if err != nil {
		uc.Logger.Error("usecase", "UserUC_Login-getUserStore", err)
		return nil, "", err
	}

	stores := []int64{}
	if len(userStores) > 0 {
		for _, store := range userStores {
			stores = append(stores, store.StoreID)
		}
	}

	user := entity.CompleteUser{
		ID:       userDetail.ID,
		Username: userDetail.Username,
		Name:     userDetail.Name,
		Role:     *role,
		StoreIDs: stores,
	}

	return &user, token, nil
}

func (uc *UserUc) validatePassword(dbPassword, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword))
	if err != nil {
		return err
	}

	return nil
}
