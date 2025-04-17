package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) Upload(c *gin.Context) {
	err := h.ProductUsecase.UploadCompleteProduct([]byte(`alright`))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusAccepted, nil)
}
