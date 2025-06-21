package entity

import "time"

const ProductTableName = "products"

type Product struct {
	ID        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	SKU       string    `json:"sku" gorm:"column:sku"`
	Barcode   string    `json:"barcode" gorm:"column:barcode"`
	StoreID   int       `json:"store_id" gorm:"column:store_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy string    `json:"updated_by" gorm:"column:updated_by"`
}

type ProductPrice struct {
	Product
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
}

type GetProductsParams struct {
	StoreID int
	Barcode string
	SKU     string
	Name    string
	Limit   int
	Offset  int
	Page    int
}

type PaginationMeta struct {
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	TotalItems   int64 `json:"total_items"`
	ItemsPerPage int   `json:"items_per_page"`
	HasNext      bool  `json:"has_next"`
	HasPrev      bool  `json:"has_prev"`
}

type PaginatedProductResponse struct {
	Data       []ProductPrice `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

func (Product) TableName() string {
	return ProductTableName
}
