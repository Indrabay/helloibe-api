package repository

import "github.com/indrabay/helloibe-api/pkg/modules/warung/entity"

func (r *PriceRepo) GetPrice(productID string) (res entity.Price, err error) {
	err = r.readDB.Where("product_id = ?", productID).First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
