package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/indrabay/helloibe-api/pkg/modules/auth/entity"
)

func (h *UserHandler) Login(c *gin.Context) {
	var params entity.LoginParam
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusUnauthorized, entity.TokenSerializer("", http.StatusUnauthorized, "param is not valid"))
		return
	}

	if err := params.Validate(); err != nil {
		c.JSON(http.StatusUnauthorized, entity.TokenSerializer("", http.StatusUnauthorized, err.Error()))
		return
	}

	token, err := h.UserUsecase.Login(params.Username, params.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.TokenSerializer("", http.StatusUnauthorized, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.TokenSerializer(token, http.StatusOK, ""))
}
