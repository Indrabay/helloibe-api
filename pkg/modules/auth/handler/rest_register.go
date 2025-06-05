package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
)

func (h *UserHandler) Register(c *gin.Context) {
	var params entity.RegisterParam
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, entity.UserSerializer(nil, http.StatusBadRequest, "param is not valid"))
		return
	}

	if err := params.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, entity.UserSerializer(nil, http.StatusBadRequest, err.Error()))
		return
	}

	user := entity.User{
		Username: params.Username,
		Password: params.Password,
		Name:     params.Name,
		Role:     params.Role,
		StoreID:  params.StoreID,
	}

	err := h.UserUsecase.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.UserSerializer(nil, http.StatusUnauthorized, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.UserSerializer(&user, http.StatusOK, ""))
}
