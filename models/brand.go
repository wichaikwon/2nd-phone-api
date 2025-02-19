package models

type Brand struct {
	BrandID uint   `gorm:"primaryKey" json:"brand_id"`
	Name    string `gorm:"not null" json:"name"`
}
