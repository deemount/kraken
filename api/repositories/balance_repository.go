// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package repositories ...
package repositories

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/deemount/kraken/api/models"
	"github.com/deemount/kraken/api/requests"
	"github.com/deemount/kraken/api/utils"
)

// BalanceRepository represents the contract between
type BalanceRepository interface {
	GetBalance() (interface{}, error)
}

// BalanceService is a struct
type BalanceService struct {
	version   int
	url       string
	uri       string
	useragent string
	key       string
	secret    string
	// balance   *models.Balance
	response models.Response
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

// GetBalance is a method
func (rs BalanceService) GetBalance() (interface{}, error) {

	var err error

	values := url.Values{}

	path := fmt.Sprintf("/%d/private/Balance", rs.version)
	url := fmt.Sprintf("%s%s", rs.url, path)
	secret, _ := base64.StdEncoding.DecodeString(rs.secret)

	// set nonce
	values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

	// create signature
	signature := utils.Signature(path, values, secret)

	// add key and signature to request headers
	headers := map[string]string{
		"API-Key":  rs.key,
		"API-Sign": signature,
	}

	q := new(requests.HTTPClient)

	body, err := q.Send(url, values, headers)
	if err != nil {
		return nil, err
	}

	// parse request
	data := rs.response // rs.balance

	err = json.Unmarshal(body, &data.Result)
	if err != nil {
		return nil, err
	}

	// check for kraken api error
	if len(data.Error) > 0 {
		return nil, fmt.Errorf("%s", data.Error)
	}

	return data.Result, err

}
