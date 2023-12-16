package model

type Product struct {
	ID           string  `json:"id" binding:"uuid4"`
	Name         string  `json:"name" binding:"omitempty,alphanum,min=1,max=255"`
	Description  string  `json:"description" binding:"omitempty,min=1,max=255"`
	VersionCount float64 `json:"version_count" binding:"omitempty,min=0,max=1000000,numeric"`
}

type ProductVersion struct {
	ID          string `json:"id" binding:"uuid4"`
	ProductID   string `json:"product_id" binding:"omitempty,uuid4"`
	Version     string `json:"version" binding:"omitempty,alphanum,min=1,max=255"`
	Description string `json:"description" binding:"omitempty,min=1,max=255"`
}
