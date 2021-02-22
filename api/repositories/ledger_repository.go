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

// LedgerRepository represents the contract between
type LedgerRepository interface {
	GetLedger(args map[string]string) (interface{}, error)
}

// LedgerService is a struct
type LedgerService struct {
	version   int
	url       string
	uri       string
	useragent string
	key       string
	secret    string
	//ledger    *models.LedgerResponse
	response models.Response
}

// NewLedgerService is a object
func NewLedgerService(version int, url, uri, useragent, key, secret string) LedgerRepository {
	return &LedgerService{
		version:   version,
		url:       url,
		uri:       uri,
		useragent: useragent,
		key:       key,
		secret:    secret,
	}
}

// GetLedger is a method
func (rs *LedgerService) GetLedger(args map[string]string) (interface{}, error) {

	var err error

	values := url.Values{}

	if value, ok := args["aclass"]; ok {
		values.Add("aclass", value)
	}
	if value, ok := args["asset"]; ok {
		values.Add("asset", value)
	}
	if value, ok := args["type"]; ok {
		values.Add("type", value)
	}
	if value, ok := args["start"]; ok {
		values.Add("start", value)
	}
	if value, ok := args["end"]; ok {
		values.Add("end", value)
	}
	if value, ok := args["ofs"]; ok {
		values.Add("ofs", value)
	}

	path := fmt.Sprintf("/%d/private/Ledgers", rs.version)
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
	data := rs.response //rs.ledger

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
