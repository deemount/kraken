package models

// Response ...
type Response struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}
