package models

import "time"

// Balance entities contains columns and register fields, types and tags
type Balance struct {
	ReportID    uint64     `gorm:"column:report_id;primary_key;unique;auto_increment" json:"report_id" example:"1"`
	EuroCash    float64    `gorm:"column:eur_cash" json:"eur" example:"12376.23"`
	CryptoAsset string     `gorm:"column:crypto_asset;varchar(100)" json:"asset" example:"XXBT"`
	CryptCash   float64    `gorm:"column:crypto_cash" json:"cryptocash" example:"1.63453"`
	CreatedAt   time.Time  `gorm:"column:created_datetime" json:"created_at" example:"2020-01-02 00:00:01+00"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at" example:"2020-01-02 00:00:01+00"`
}
