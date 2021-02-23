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

// TradeBalanceRepository represents the contract between
type TradeBalanceRepository interface {
	GetTradeBalance(args map[string]string) (interface{}, error)
}

// TradeBalanceService is a struct
type TradeBalanceService struct {
	version   int
	url       string
	uri       string
	useragent string
	key       string
	secret    string
	// tradebalance *models.TradeBalance
	response models.Response
}

// NewTradeBalanceService is a object
func NewTradeBalanceService(version int, url, uri, useragent, key, secret string) TradeBalanceRepository {
	return &TradeBalanceService{
		version:   version,
		url:       url,
		uri:       uri,
		useragent: useragent,
		key:       key,
		secret:    secret,
	}
}

// GetTradeBalance returns trade balance info
func (rs TradeBalanceService) GetTradeBalance(args map[string]string) (interface{}, error) {

	var err error

	values := url.Values{}
	if value, ok := args["aclass"]; ok {
		values.Add("aclass", value)
	}
	if value, ok := args["asset"]; ok {
		values.Add("asset", value)
	}

	path := fmt.Sprintf("/%d/private/TradeBalance", rs.version)
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
	data := rs.response // rs.tradebalance

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
