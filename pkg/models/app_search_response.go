package models

// AppSearchResponse represents app search response
type AppSearchResponse struct {
	Count  int           `json:"Count"  yaml:"Count" binding:"required"`
	Offset int           `json:"Offset"  yaml:"Offset"`
	Limit  int           `json:"Limit"  yaml:"Limit"`
	Data   []AppMetadata `json:"Data"  yaml:"Data" binding:"required"`
}
