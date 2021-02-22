// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package models ...
package models

// Balance entities contains columns and register fields, types and tags
// type Balance struct {
// 	ReportID    uint64     `gorm:"column:report_id;primary_key;unique;auto_increment" json:"report_id" example:"1"`
// 	EuroCash    float64    `gorm:"column:eur_cash" json:"eur" example:"12376.23"`
// 	CryptoAsset string     `gorm:"column:crypto_asset;varchar(100)" json:"asset" example:"XXBT"`
// 	CryptCash   float64    `gorm:"column:crypto_cash" json:"cryptocash" example:"1.63453"`
// 	CreatedAt   time.Time  `gorm:"column:created_datetime" json:"created_at" example:"2020-01-02 00:00:01+00"`
// 	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at" example:"2020-01-02 00:00:01+00"`
// }

// Balance ...
type Balance struct {
	ADA  float64 `json:"ADA,string"`
	AAVE float64 `json:"AAVE,string"`
	BCH  float64 `json:"BCH,string"`
	DASH float64 `json:"DASH,string"`
	EOS  float64 `json:"EOS,string"`
	GNO  float64 `json:"GNO,string"`
	QTUM float64 `json:"QTUM,string"`
	KFEE float64 `json:"KFEE,string"`
	USDT float64 `json:"USDT,string"`
	XDAO float64 `json:"XDAO,string"`
	XETC float64 `json:"XETC,string"`
	XETH float64 `json:"XETH,string"`
	XICN float64 `json:"XICN,string"`
	XLTC float64 `json:"XLTC,string"`
	XMLN float64 `json:"XMLN,string"`
	XNMC float64 `json:"XNMC,string"`
	XREP float64 `json:"XREP,string"`
	XXBT float64 `json:"XXBT,string"`
	XXDG float64 `json:"XXDG,string"`
	XXLM float64 `json:"XXLM,string"`
	XXMR float64 `json:"XXMR,string"`
	XXRP float64 `json:"XXRP,string"`
	XTZ  float64 `json:"XTZ,string"`
	XXVN float64 `json:"XXVN,string"`
	XZEC float64 `json:"XZEC,string"`
	ZCAD float64 `json:"ZCAD,string"`
	ZEUR float64 `json:"ZEUR,string"`
	ZGBP float64 `json:"ZGBP,string"`
	ZJPY float64 `json:"ZJPY,string"`
	ZKRW float64 `json:"ZKRW,string"`
	ZUSD float64 `json:"ZUSD,string"`
	TRX  float64 `json:"TRX,string"`
}
