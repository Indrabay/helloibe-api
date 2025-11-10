package handler

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	source, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to open file"})
		return
	}

	defer source.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, source)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read file"})
		return
	}

	err = h.ProductUsecase.UploadProductPrice(c, buf.Bytes(), file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "file uploaded successfully"})
}
