package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUsername(username string) (*entity.User, error)
	GetRole(roleID int) (*entity.Role, error)
	Insert(*entity.User) error
}

type UserRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewUserRepository(writeDB, readDB *gorm.DB) *UserRepo {
	return &UserRepo{
		writeDB: writeDB,
		readDB:  readDB,
	}
}
