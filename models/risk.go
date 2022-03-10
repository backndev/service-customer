package models

type RiskProfile struct {
	Id           uint    `json:"id"`
	MmPercent    float32 `json:"mm-percent"`
	BondPercent  float32 `json:"bond-percent"`
	StockPercent float32 `json:"stock-percent"`
	UserId       uint    `json:"user_id"`
	User         User    `json:"user" gorm:"foreignKey:UserId"`
}
