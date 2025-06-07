package repository

import "github.com/indrabay/helloibe-api/pkg/modules/warung/entity"

func (r *PriceRepo) GetPrices(productIDs []string) (res []entity.Price, err error) {
	err = r.readDB.Where("product_id IN ?", productIDs).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
