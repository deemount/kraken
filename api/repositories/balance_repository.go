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

// BalanceRepository represents the contract between
type BalanceRepository interface {
	FindBalanceByCurrency(values url.Values) (*models.Balance, error)
}

// BalanceService is a struct
type BalanceService struct {
	version   string
	url       string
	uri       string
	useragent string
	key       string
	secret    string
	balance   *models.Balance
}

// NewBalanceService is a object
func NewBalanceService(url, uri, useragent, key, secret string) BalanceRepository {
	return &BalanceService{
		url:       url,
		uri:       uri,
		useragent: useragent,
		key:       key,
		secret:    secret,
	}
}

// FindBalanceByCurrency is a method
func (rs *BalanceService) FindBalanceByCurrency(values url.Values) (*models.Balance, error) {

	var err error

	path := fmt.Sprintf("%s/private/Balance", rs.version)
	url := fmt.Sprintf("%s%s", rs.url, path)
	secret, _ := base64.StdEncoding.DecodeString(rs.secret)

	// set nonce
	values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

	// create signature
	signature := utils.Signature(path, values, secret)

	// add token to request headers
	headers := map[string]string{
		"API-Key":  rs.key,
		"API-Sign": signature,
	}

	q := new(requests.HTTPClient)

	body, err := q.Send(url, values, headers)
	if err != nil {
		log.Print(err)
	}

	err = json.Unmarshal(body, &rs.balance)
	return rs.balance, err

}
