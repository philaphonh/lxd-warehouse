package model

import "time"

// Import type
// Use for sending data to user only
type Import struct {
	ID         int    `json:"id"`
	Supplier   string `json:"supplier"`
	User       string `json:"user"`
	ImportTime string `json:"importTime"`
	Status     string `json:"status"`
}

// ImportRaw type
type ImportRaw struct {
	ID         int       `json:"id"`
	SupplierID string    `json:"supplier_id"`
	UserID     string    `json:"user_id"`
	ImportTime time.Time `json:"importTime"`
	Status     string    `json:"status"`
}

// ImportDetail type
type ImportDetail struct {
	ID       int    `json:"id"`
	ImportID int    `json:"importId"`
	Product  string `json:"product"`
	Qty      int    `json:"qty"`
}
