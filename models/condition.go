package models

type Condition struct {
	ConditionID      uint   `gorm:"primaryKey" json:"condition_id"`
	PhoneID          uint   `gorm:"not null" json:"phone_id"`
	Scratches        bool   `json:"scratches"`
	Bend             bool   `json:"bend"`
	BackGlassCrack   bool   `json:"back_glass_cracked"`
	ScreenStatus     string `json:"screen_condition"`
	BatteryHealth    string `json:"battery_health"`
	Accessories      string `json:"accessories"`
	FunctionalIssues string `json:"functional_issues"`
}
