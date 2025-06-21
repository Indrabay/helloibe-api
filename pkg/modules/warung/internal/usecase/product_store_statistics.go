package usecase

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

// GetProductCount retrieves only the product count for a store
func (uc *ProductUc) GetProductCount(storeID int64) (int64, error) {
	return uc.ProductRepo.GetProductCountByStore(storeID)
}

// GetPriceCount retrieves only the price count for a store
func (uc *ProductUc) GetPriceCount(storeID int64) (int64, error) {
	return uc.ProductRepo.GetPriceCountByStore(storeID)
}

// GetStoreStatistics retrieves complete store statistics
func (uc *ProductUc) GetStoreStatistics(storeID int64) (entity.StoreStatistics, error) {
	return uc.ProductRepo.GetStoreStatistics(storeID)
}
