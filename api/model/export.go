package model

// Export type
type Export struct {
	ID           int    `json:"id"`
	Distributor  string `json:"distributor"`
	User         string `json:"user"`
	ExportTime   string `json:"exportTime"`
	ExportStatus string `json:"exportStatus"`
}

// ExportDetail type
type ExportDetail struct {
	ID       int    `json:"id"`
	ExportID int    `json:"exportId"`
	Product  string `json:"product"`
	Qty      int    `json:"qty"`
}
