package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/config"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/internal/repository"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/internal/usecase"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductHandler(cfg config.WarungConfig) *ProductHandler {
	productRepo := repository.NewProductRepository(cfg.WriteDB, cfg.ReadDB)
	priceRepo := repository.NewPriceRepository(cfg.WriteDB, cfg.ReadDB)
	productUc := usecase.NewProductUsecase(productRepo, priceRepo)
	return &ProductHandler{
		ProductUsecase: productUc,
	}
}

func (h *ProductHandler) MountProduct(group *gin.RouterGroup) {
	group.Group("/products")
	group.POST("/uploads", h.Upload)
}
