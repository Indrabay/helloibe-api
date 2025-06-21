package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/config"
	"github.com/indrabay/helloibe-api/pkg/middleware"
	"github.com/indrabay/helloibe-api/pkg/modules/auth/handler"
	warungHandler "github.com/indrabay/helloibe-api/pkg/modules/warung/handler"
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
	routes.Use(middleware.CORSMiddleware())
	userGroup := routes.Group("/users")
	userHandler.MountUser(userGroup)

	warungConfig := config.WarungConfig{
		WriteDB: writeDB,
		ReadDB:  readDB,
	}

	warungHandler := warungHandler.NewProductHandler(warungConfig)

	warungGroup := routes.Group("/warung")
	warungGroup.Use(middleware.Auth())
	warungHandler.MountProduct(warungGroup)

	return routes
}
