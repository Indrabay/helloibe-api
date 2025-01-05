package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s?parseTime=True", dsn)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)

	return db, nil
}
