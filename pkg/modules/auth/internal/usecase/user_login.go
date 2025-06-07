package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"github.com/indrabay/helloibe-api/utils"
	"github.com/indrabay/helloibe-api/utils/logger"
	"go.uber.org/zap"
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
		logger.Error(err.Error(),
			zap.String("method", "UserUc_Login"),
			zap.String("step", "GetByUsername"),
			zap.String("username", username),
		)
		return nil, "", utils.ErrUserNotFound
	}

	err = uc.validatePassword(userDetail.Password, password)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_Login"),
			zap.String("step", "validatePassword"),
		)
		return nil, "", utils.ErrUserNotFound
	}

	role, err := uc.UserRepo.GetRole(userDetail.Role)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_Login"),
			zap.String("step", "GetRole"),
			zap.Int("role", userDetail.Role),
		)
		return nil, "", err
	}

	userStores, err := uc.UserRepo.GetUserStores(userDetail.ID)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_Login"),
			zap.String("step", "GetUserStores"),
			zap.Int64("user_id", userDetail.ID),
		)
		return nil, "", err
	}

	stores := []int64{}
	if len(userStores) > 0 {
		for _, store := range userStores {
			stores = append(stores, store.StoreID)
		}
	}

	claims := utils.JWTClaim{
		Username: userDetail.Username,
		Name:     userDetail.Name,
		Role:     userDetail.Role,
		Stores:   stores,
	}

	token, err = uc.JWTUtils.CreateToken(claims)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_Login"),
			zap.String("step", "CreateToken"),
		)
		return nil, "", err
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
