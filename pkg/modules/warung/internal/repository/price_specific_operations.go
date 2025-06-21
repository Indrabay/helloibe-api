package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

// GetPriceByID retrieves a single price by its ID
func (r *PriceRepo) GetPriceByID(priceID string) (price entity.Price, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("id = ?", priceID).First(&price).Error
	return price, err
}

// GetPricesByProductID retrieves all prices for a specific product
func (r *PriceRepo) GetPricesByProductID(productID string) (prices []entity.Price, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("product_id = ?", productID).Find(&prices).Error
	return prices, err
}

// GetPriceByProductID retrieves a single price for a specific product
func (r *PriceRepo) GetPriceByProductID(productID string) (price entity.Price, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("product_id = ?", productID).First(&price).Error
	return price, err
}

// DeletePriceByID deletes a single price by its ID
func (r *PriceRepo) DeletePriceByID(priceID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Where("id = ?", priceID).Delete(&entity.Price{}).Error
}

// DeletePriceByProductID deletes all prices for a specific product
func (r *PriceRepo) DeletePriceByProductID(productID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Where("product_id = ?", productID).Delete(&entity.Price{}).Error
}

// UpdatePrice updates a single price
func (r *PriceRepo) UpdatePrice(price entity.Price) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Save(&price).Error
}

// GetPricesBySellingPriceRange retrieves prices within a selling price range
func (r *PriceRepo) GetPricesBySellingPriceRange(minPrice, maxPrice float64) (prices []entity.Price, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("selling_price BETWEEN ? AND ?", minPrice, maxPrice).Find(&prices).Error
	return prices, err
}
