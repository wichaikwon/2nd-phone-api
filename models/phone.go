package models

type Phone struct {
	PhoneID       uint    `gorm:"primaryKey" json:"phone_id"`
	BrandID       uint    `gorm:"not null" json:"brand_id"`
	ModelID       uint    `gorm:"not null" json:"model_id"`
	Storage       string  `json:"storage"`
	Warranty      string  `json:"warranty_status"`
	Origin        string  `json:"origin"`
	OriginalPrice float64 `json:"original_price"`
	CurrentPrice  float64 `json:"current_price"`
}
