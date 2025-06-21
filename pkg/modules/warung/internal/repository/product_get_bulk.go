package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (r *ProductRepo) GetProducts(params entity.GetProductsParams) (res []entity.Product, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var bindParam []any
	query := "SELECT * FROM products WHERE store_id = ?"
	bindParam = append(bindParam, params.StoreID)

	if params.Barcode != "" {
		query += " AND barcode = ?"
		bindParam = append(bindParam, params.Barcode)
	}

	if params.SKU != "" {
		query += " AND sku = ?"
		bindParam = append(bindParam, params.SKU)
	}

	if params.Name != "" {
		query += " AND name like ?"
		bindParam = append(bindParam, "%"+params.Name+"%")
	}

	// Add ORDER BY for consistent pagination
	query += " ORDER BY created_at DESC"

	// Add pagination
	if params.Limit > 0 {
		query += " LIMIT ?"
		bindParam = append(bindParam, params.Limit)
	}

	if params.Offset > 0 {
		query += " OFFSET ?"
		bindParam = append(bindParam, params.Offset)
	}

	err = r.readDB.Raw(query, bindParam...).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetProductsCount returns the total count of products matching the search criteria
func (r *ProductRepo) GetProductsCount(params entity.GetProductsParams) (count int64, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var bindParam []any
	query := "SELECT COUNT(*) FROM products WHERE store_id = ?"
	bindParam = append(bindParam, params.StoreID)

	if params.Barcode != "" {
		query += " AND barcode = ?"
		bindParam = append(bindParam, params.Barcode)
	}

	if params.SKU != "" {
		query += " AND sku = ?"
		bindParam = append(bindParam, params.SKU)
	}

	if params.Name != "" {
		query += " AND name like ?"
		bindParam = append(bindParam, "%"+params.Name+"%")
	}

	err = r.readDB.Raw(query, bindParam...).Scan(&count).Error
	return count, err
}
