package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	source, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	defer source.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, source)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	username, exists := c.Get("username")
	fmt.Println("haha", username, exists)

	err = h.ProductUsecase.UploadProductPrice(c, buf.Bytes())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusAccepted, nil)
}
