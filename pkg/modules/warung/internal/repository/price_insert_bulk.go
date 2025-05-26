package repository

import (
	"github.com/indrabay/helloibe-api/pkg/modules/warung/entity"
	"gorm.io/gorm/clause"
)

func (p *PriceRepo) InsertBulk(prices []entity.Price) error {
	err := p.writeDB.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"purchase_price", "selling_price"}),
	}).CreateInBatches(&prices, BatchSize).Error

	return err
}
