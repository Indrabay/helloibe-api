package usecase

import (
	"errors"

	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"github.com/indrabay/helloibe-api/utils/logger"
	"go.uber.org/zap"
)

// AddStoreToUser adds a store to a user, with validation for one store per user (except super admin)
func (uc *UserUc) AddStoreToUser(userID int64, storeID int64, roleID int) error {
	// Super admin can have multiple stores
	if roleID == entity.SuperAdminRoleID {
		userStore := entity.UserStore{
			UserID:  userID,
			StoreID: storeID,
		}
		return uc.UserRepo.InsertUserStore(userStore)
	}

	// For non-super-admin users, check if they already have a store
	storeCount, err := uc.UserRepo.GetUserStoreCount(userID)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_AddStoreToUser"),
			zap.String("step", "GetUserStoreCount"),
			zap.Int64("userID", userID),
		)
		return err
	}

	if storeCount > 0 {
		return errors.New("user can only have one store (except super admin)")
	}

	userStore := entity.UserStore{
		UserID:  userID,
		StoreID: storeID,
	}

	err = uc.UserRepo.InsertUserStore(userStore)
	if err != nil {
		logger.Error(err.Error(),
			zap.String("method", "UserUc_AddStoreToUser"),
			zap.String("step", "InsertUserStore"),
			zap.Any("userStore", userStore),
		)
		return err
	}

	return nil
}

