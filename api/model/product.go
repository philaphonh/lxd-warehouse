package model

// Product type
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	Image       string  `json:"imageUrl"`
}

// ProductRaw type
type ProductRaw struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	BrandID     int     `json:"brand_id"`
	TypeID      int     `json:"type_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	Image       string  `json:"imageUrl"`
}

// ImportedProduct type
type ImportedProduct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	BrandID     int     `json:"brand_id"`
	TypeID      int     `json:"type_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	Quantity    string  `json:"quantity"`
	Image       string  `json:"imageUrl"`
}

// ImportedProductDetail type
type ImportedProductDetail struct {
	ProductID  int `json:"productId"`
	ProductQty int `json:"productQty"`
}

// ExportedProduct type
type ExportedProduct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	BrandID     int     `json:"brand_id"`
	TypeID      int     `json:"type_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	Quantity    string  `json:"quantity"`
	Image       string  `json:"imageUrl"`
}

// ExportedProductDetail type
type ExportedProductDetail struct {
	ProductID  int `json:"productId"`
	ProductQty int `json:"productQty"`
}
