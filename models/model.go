package models

type Model struct {
	ModelID     uint   `gorm:"primaryKey" json:"model_id"`
	BrandID     uint   `gorm:"not null" json:"brand_id"`
	Name        string `gorm:"not null" json:"name"`
	ReleaseYear int    `json:"release_year"`
}
