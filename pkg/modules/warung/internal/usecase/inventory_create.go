package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (uc *InventoryUc) Create(ctx *gin.Context, inventory *entity.Inventory) error {
	username, exist := ctx.Get("username")
	if !exist {
		return nil
	}

	// Generate ID if not provided
	if inventory.ID == "" {
		inventory.ID = uuid.NewString()
	}

	// Set audit fields
	inventory.CreatedBy = username.(string)
	inventory.UpdatedBy = username.(string)

	// Insert inventory
	err := uc.InventoryRepo.Insert(inventory)
	if err != nil {
		return err
	}

	// Initialize quantity to 0 (will be updated when details are added)
	inventory.Quantity = 0
	err = uc.InventoryRepo.Update(inventory)
	if err != nil {
		return err
	}

	return nil
}
