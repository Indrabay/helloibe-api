package repository

import "github.com/indrabay/helloibe-api/pkg/modules/warung/entity"

func (p *ProductRepo) GetProductsByBarcode(barcodes []string) (res []entity.Product, err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	err = p.readDB.Where("barcode IN ?", barcodes).Find(&res).Error

	return res, err
}
