// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package models ...
package models

// TradeBalance ...
type TradeBalance struct {
	EquivalentBalance         float64 `json:"eb,string"`
	TradeBalance              float64 `json:"tb,string"`
	MarginOP                  float64 `json:"m,string"`
	UnrealizedNetProfitLossOP float64 `json:"n,string"`
	CostBasisOP               float64 `json:"c,string"`
	CurrentValuationOP        float64 `json:"v,string"`
	Equity                    float64 `json:"e,string"`
	FreeMargin                float64 `json:"mf,string"`
	MarginLevel               float64 `json:"ml,string"`
}
