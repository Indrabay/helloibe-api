package repository

import "github.com/indrabay/helloibe-api/pkg/modules/warung/entity"

func (p *ProductRepo) GetProductsByBarcode(barcodes []string) (res []entity.Product, err error) {
	err = p.readDB.Where("barcode IN ?", barcodes).Find(&res).Error

	return res, err
}
