package models

import "math/big"

// Ledger ...
type Ledger struct {
}

// LedgerResponse represents an associative array of ledgers infos
type LedgerResponse struct {
	Ledger map[string]LedgerInfo `json:"ledger"`
}

// LedgerInfo represents the ledger informations
type LedgerInfo struct {
	RefID   string    `json:"refid"`
	Time    float64   `json:"time"`
	Type    string    `json:"type"`
	Aclass  string    `json:"aclass"`
	Asset   string    `json:"asset"`
	Amount  big.Float `json:"amount"`
	Fee     big.Float `json:"fee"`
	Balance big.Float `json:"balance"`
}
