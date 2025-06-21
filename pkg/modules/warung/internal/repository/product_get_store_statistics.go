package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

// GetProductCountByStore counts only the number of products for a specific store
func (r *ProductRepo) GetProductCountByStore(storeID int64) (count int64, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Model(&entity.Product{}).Where("store_id = ?", storeID).Count(&count).Error
	return count, err
}

// GetPriceCountByStore counts only the number of prices for products in a specific store
func (r *ProductRepo) GetPriceCountByStore(storeID int64) (count int64, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Model(&entity.Price{}).
		Joins("JOIN products ON prices.product_id = products.id").
		Where("products.store_id = ?", storeID).
		Count(&count).Error
	return count, err
}

// GetStoreStatistics combines product and price counts for a store
func (r *ProductRepo) GetStoreStatistics(storeID int64) (res entity.StoreStatistics, err error) {
	productCount, err := r.GetProductCountByStore(storeID)
	if err != nil {
		return res, err
	}

	priceCount, err := r.GetPriceCountByStore(storeID)
	if err != nil {
		return res, err
	}

	res = entity.StoreStatistics{
		StoreID:      storeID,
		ProductCount: productCount,
		PriceCount:   priceCount,
	}

	return res, nil
}
