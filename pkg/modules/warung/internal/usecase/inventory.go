package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/internal/repository"
)

type InventoryUsecase interface {
	Create(ctx *gin.Context, inventory *entity.Inventory) error
	Update(ctx *gin.Context, id string, inventory *entity.Inventory) error
	GetByID(id string) (entity.Inventory, error)
	GetByFilter(params entity.GetInventoriesParams) (entity.PaginatedInventoryResponse, error)
	Upload(ctx *gin.Context, value []byte, filename string) error
}

type InventoryUc struct {
	InventoryRepo       repository.InventoryRepository
	InventoryDetailRepo repository.InventoryDetailRepository
	ProductRepo         repository.ProductRepository
}

func NewInventoryUsecase(
	inventoryRepo repository.InventoryRepository,
	inventoryDetailRepo repository.InventoryDetailRepository,
	productRepo repository.ProductRepository,
) *InventoryUc {
	return &InventoryUc{
		InventoryRepo:       inventoryRepo,
		InventoryDetailRepo: inventoryDetailRepo,
		ProductRepo:         productRepo,
	}
}
