package entity

type StoreStatistics struct {
	StoreID      int64 `json:"store_id"`
	ProductCount int64 `json:"product_count"`
	PriceCount   int64 `json:"price_count"`
}
