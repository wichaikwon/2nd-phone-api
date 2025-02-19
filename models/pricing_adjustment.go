package models

type PricingAdjustment struct {
	AdjustmentID    uint    `gorm:"primaryKey" json:"adjustment_id"`
	PhoneID         uint    `gorm:"not null" json:"phone_id"`
	ConditionID     uint    `gorm:"not null" json:"condition_id"`
	PriceAdjustment float64 `json:"price_adjustment"`
}
