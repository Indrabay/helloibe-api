package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"github.com/indrabay/helloibe-api/pkg/modules/auth/internal/repository"
	"github.com/indrabay/helloibe-api/utils"
)

type UserUsecase interface {
	// Login authenticates a user with username and password, returns JWT token if successful
	Login(username, password string) (*entity.CompleteUser, string, error)
	// Register creates a new user account in the system
	Register(user *entity.User) error
}

// TODO: add metrics
type UserUc struct {
	UserRepo repository.UserRepository
	JWTUtils utils.JWT
}

func NewUserUsecase(userRepo repository.UserRepository, jwtUtils utils.JWT) *UserUc {
	return &UserUc{
		UserRepo: userRepo,
		JWTUtils: jwtUtils,
	}
}
