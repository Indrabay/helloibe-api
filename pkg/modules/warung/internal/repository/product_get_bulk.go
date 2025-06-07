package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
)

func (r *ProductRepo) GetProducts(params entity.GetProductsParams) (res []entity.Product, err error) {
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

	err = r.readDB.Raw(query, bindParam...).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
