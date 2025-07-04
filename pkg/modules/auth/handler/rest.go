package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/config"
	"github.com/indrabay/helloibe-api/pkg/middleware"
	"github.com/indrabay/helloibe-api/pkg/modules/auth/internal/repository"
	"github.com/indrabay/helloibe-api/pkg/modules/auth/internal/usecase"
	"github.com/indrabay/helloibe-api/utils"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(cfg config.UserConfig) *UserHandler {
	userRepo := repository.NewUserRepository(cfg.WriteDB, cfg.ReadDB)
	jwtUtils := utils.NewJWT(utils.Config.SigningKey)

	userUsecase := usecase.NewUserUsecase(userRepo, jwtUtils)
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) MountUser(group *gin.RouterGroup) {
	group.POST("/login", h.Login)
	group.Use(middleware.Auth())
	group.POST("/register", h.Register)
}
