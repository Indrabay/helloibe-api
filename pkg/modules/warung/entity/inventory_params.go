package entity

type GetInventoriesParams struct {
	ProductID string
	StoreID   int
	Limit     int
	Offset    int
	Page      int
}

type PaginatedInventoryResponse struct {
	Data       []Inventory    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}
