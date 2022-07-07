package models

// AppSearchResponse represents app search response
type AppSearchResponse struct {
	Count  int           `json:"count"  yaml:"count" binding:"required"`
	Offset int           `json:"offset"  yaml:"offset"`
	Limit  int           `json:"limit"  yaml:"limit"`
	Data   []AppMetadata `json:"data"  yaml:"data" binding:"required"`
}
