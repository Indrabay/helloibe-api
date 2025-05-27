package config

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type WarungConfig struct {
	WriteDB   *gorm.DB
	ReadDB    *gorm.DB
	ZapLogger *zap.SugaredLogger
}
