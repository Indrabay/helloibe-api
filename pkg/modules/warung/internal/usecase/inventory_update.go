package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (uc *InventoryUc) Update(ctx *gin.Context, id string, inventory *entity.Inventory) error {
	username, exist := ctx.Get("username")
	if !exist {
		return nil
	}

	// Get existing inventory
	existing, err := uc.InventoryRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Update fields
	existing.ProductID = inventory.ProductID
	existing.UpdatedBy = username.(string)

	// Update inventory
	err = uc.InventoryRepo.Update(&existing)
	if err != nil {
		return err
	}

	// Update quantity from details
	err = uc.InventoryRepo.UpdateInventoryQuantity(id)
	if err != nil {
		return err
	}

	return nil
}
