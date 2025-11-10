package usecase

import (
	"math"

	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (uc *InventoryUc) GetByID(id string) (inventory entity.Inventory, err error) {
	return uc.InventoryRepo.GetByID(id)
}

func (uc *InventoryUc) GetByFilter(params entity.GetInventoriesParams) (response entity.PaginatedInventoryResponse, err error) {
	// Get total count for pagination metadata
	totalItems, err := uc.InventoryRepo.GetCount(params)
	if err != nil {
		return response, err
	}

	// Get paginated inventories
	inventories, err := uc.InventoryRepo.GetByFilter(params)
	if err != nil {
		return response, err
	}

	// Calculate pagination metadata
	totalPages := int(math.Ceil(float64(totalItems) / float64(params.Limit)))
	if totalPages == 0 {
		totalPages = 1
	}

	hasNext := params.Page < totalPages
	hasPrev := params.Page > 1

	response = entity.PaginatedInventoryResponse{
		Data: inventories,
		Pagination: entity.PaginationMeta{
			CurrentPage:  params.Page,
			TotalPages:   totalPages,
			TotalItems:   totalItems,
			ItemsPerPage: params.Limit,
			HasNext:      hasNext,
			HasPrev:      hasPrev,
		},
	}

	return response, nil
}
