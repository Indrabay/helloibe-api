package repository

import "github.com/indrabay/helloibe-api/pkg/modules/warung/entity"

func (r *ProductRepo) GetProductByBarcodeStore(barcode string, storeID int64) (res entity.Product, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.readDB.Where("barcode = ? AND store_id = ?", barcode, storeID).First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
