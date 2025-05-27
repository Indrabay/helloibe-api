package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm/clause"
)

const BatchSize = 50

func (p *ProductRepo) InsertBulk(products []entity.Product) error {
	err := p.writeDB.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"name", "sku", "updated_by"}),
	}).CreateInBatches(&products, BatchSize).Error

	return err
}
