package models

type RiskProfile struct {
	Id           uint    `json:"id"`
	MmPercent    float32 `json:"mm_percent"`
	BondPercent  float32 `json:"bond_percent"`
	StockPercent float32 `json:"stock_percent"`
	UserId       uint    `json:"user_id"`
	User         User    `json:"user" gorm:"foreignKey:UserId"`
}
