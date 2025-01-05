package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/config"
	"github.com/indrabay/helloibe-api/pkg/modules/auth/handler"
	"github.com/indrabay/helloibe-api/utils"
	"go.uber.org/zap"
)

func StartServer() *gin.Engine {
	utils.LoadEnv()
	writeDB, err := utils.NewDatabase(utils.Config.DSN)
	if err != nil {
		panic(err)
	}

	readDB, err := utils.NewDatabase(utils.Config.DSN)
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	sugarLogger := logger.Sugar()

	userConfig := config.UserConfig{
		WriteDB:   writeDB,
		ReadDB:    readDB,
		ZapLogger: sugarLogger,
	}

	userHandler := handler.NewUserHandler(userConfig)

	routes := gin.Default()
	userGroup := routes.Group("/users")
	userHandler.MountUser(userGroup)

	return routes
}
