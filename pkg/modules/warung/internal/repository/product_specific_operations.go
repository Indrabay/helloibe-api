package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

// GetProductByID retrieves a single product by its ID
func (r *ProductRepo) GetProductByID(productID string) (product entity.Product, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("id = ?", productID).First(&product).Error
	return product, err
}

// GetProductsByStore retrieves all products for a specific store
func (r *ProductRepo) GetProductsByStore(storeID int64) (products []entity.Product, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("store_id = ?", storeID).Find(&products).Error
	return products, err
}

// GetProductBySKU retrieves a product by its SKU
func (r *ProductRepo) GetProductBySKU(sku string, storeID int64) (product entity.Product, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("sku = ? AND store_id = ?", sku, storeID).First(&product).Error
	return product, err
}

// GetProductByName retrieves products by name (partial match)
func (r *ProductRepo) GetProductByName(name string, storeID int64) (products []entity.Product, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("name LIKE ? AND store_id = ?", "%"+name+"%", storeID).Find(&products).Error
	return products, err
}

// DeleteProductByID deletes a single product by its ID
func (r *ProductRepo) DeleteProductByID(productID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Where("id = ?", productID).Delete(&entity.Product{}).Error
}

// UpdateProduct updates a single product
func (r *ProductRepo) UpdateProduct(product entity.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.writeDB.Save(&product).Error
}
