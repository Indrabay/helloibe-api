package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
	"github.com/indrabay/helloibe-api/utils"
)

func (h *UserHandler) Login(c *gin.Context) {
	var params entity.LoginParam
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusUnauthorized, entity.LoginResponseSerializer(nil, "", http.StatusUnauthorized, "param is not valid"))
		return
	}

	if err := params.Validate(); err != nil {
		c.JSON(http.StatusUnauthorized, entity.LoginResponseSerializer(nil, "", http.StatusUnauthorized, err.Error()))
		return
	}

	user, token, err := h.UserUsecase.Login(params.Username, params.Password)
	if err != nil {
		if err == utils.ErrUserNotFound {
			c.JSON(http.StatusUnauthorized, entity.LoginResponseSerializer(nil, "", http.StatusUnauthorized, err.Error()))
		}
		c.JSON(http.StatusInternalServerError, entity.LoginResponseSerializer(nil, "", http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.LoginResponseSerializer(user, token, http.StatusOK, ""))
}
