package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertBulk(products []entity.Product) error
}

type ProductRepo struct {
	writeDB *gorm.DB
	readDB  *gorm.DB
}

func NewProductRepository(writeDB, readDB *gorm.DB) *ProductRepo {
	return &ProductRepo{
		writeDB: writeDB,
		readDB:  readDB,
	}
}
