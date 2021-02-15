package repositories

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/deemount/kraken/api/models"
	"github.com/deemount/kraken/api/requests"
	"github.com/deemount/kraken/api/utils"
)

// KrakenResponse wraps the Kraken API JSON response
type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

// BalanceRepository represents the contract between
type BalanceRepository interface {
	FindBalanceByCurrency(values url.Values) (interface{}, error)
}

// BalanceService is a struct
type BalanceService struct {
	version   int
	url       string
	uri       string
	useragent string
	key       string
	secret    string
	balance   *models.Balance
}

// NewBalanceService is a object
func NewBalanceService(version int, url, uri, useragent, key, secret string) BalanceRepository {
	return &BalanceService{
		version:   version,
		url:       url,
		uri:       uri,
		useragent: useragent,
		key:       key,
		secret:    secret,
	}
}

// FindBalanceByCurrency is a method
func (rs *BalanceService) FindBalanceByCurrency(values url.Values) (interface{}, error) {

	var err error

	path := fmt.Sprintf("/%d/private/Balance", rs.version)
	url := fmt.Sprintf("%s%s", rs.url, path)
	secret, _ := base64.StdEncoding.DecodeString(rs.secret)

	// set nonce
	values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

	// create signature
	signature := utils.Signature(path, values, secret)

	log.Printf("Signature: %s", signature)

	// add token to request headers
	headers := map[string]string{
		"API-Key":  rs.key,
		"API-Sign": signature,
	}

	q := new(requests.HTTPClient)

	body, err := q.Send(url, values, headers)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! #3 (%s)", err.Error())
	}

	// Parse request
	var jsonData KrakenResponse

	typ := models.Balance{}

	// Set the KrakenResponse.Result to typ so `json.Unmarshal` will
	// unmarshal it into given type, instead of `interface{}`.
	jsonData.Result = typ

	err = json.Unmarshal(body, &jsonData.Result)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! #6 (%s)", err.Error())
	}

	// Check for Kraken API error
	if len(jsonData.Error) > 0 {
		return nil, fmt.Errorf("Could not execute request! #7 (%s)", jsonData.Error)
	}

	log.Printf("%+v", &jsonData)

	return jsonData.Result, err

}
