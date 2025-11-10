package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/config"
	"github.com/indrabay/helloibe-api/pkg/middleware"
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
	group.Use(middleware.ValidateStore())
	group.POST("/:store_id/products/upload", h.Upload)
	group.GET("/:store_id/products/:sku", h.SingleProduct)
	group.GET("/:store_id/products/prices", h.GetProducts)
	group.GET("/:store_id/products/prices/all", h.GetAllProductPrices)
	group.GET("/:store_id", h.GetStoreStatistics)
}

type InventoryHandler struct {
	InventoryUsecase usecase.InventoryUsecase
}

func NewInventoryHandler(cfg config.WarungConfig) *InventoryHandler {
	inventoryRepo := repository.NewInventoryRepository(cfg.WriteDB, cfg.ReadDB)
	inventoryDetailRepo := repository.NewInventoryDetailRepository(cfg.WriteDB, cfg.ReadDB)
	productRepo := repository.NewProductRepository(cfg.WriteDB, cfg.ReadDB)
	inventoryUc := usecase.NewInventoryUsecase(inventoryRepo, inventoryDetailRepo, productRepo)
	return &InventoryHandler{
		InventoryUsecase: inventoryUc,
	}
}

func (h *InventoryHandler) MountInventory(group *gin.RouterGroup) {
	group.Use(middleware.ValidateStore())
	group.POST("/:store_id/inventories", h.Create)
	group.PATCH("/:store_id/inventories/:id", h.Update)
	group.GET("/:store_id/inventories/:id", h.GetByID)
	group.GET("/:store_id/inventories", h.GetByFilter)
	group.POST("/:store_id/inventories/upload", h.Upload)
}
