// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package repositories ...
package repositories

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/deemount/kraken/api/utils"
)

func TestGetTradeBalance(t *testing.T) {

	t.Run("returns tradebalance", func(t *testing.T) {

		values := url.Values{}
		secret, _ := base64.StdEncoding.DecodeString(os.Getenv("API_KRAKEN_SECRET"))
		values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))
		headers := map[string]string{
			"API-Key":  os.Getenv("API_KRAKEN_KEY"),
			"API-Sign": utils.Signature("/0/private/TradeBalance", values, secret),
		}

		request, _ := http.NewRequest(http.MethodPost, "https://api.kraken.com/0/private/TradeBalance", strings.NewReader(values.Encode()))
		request.Header.Add("User-Agent", "GoKrakenBot/1.0")
		for key, value := range headers {
			request.Header.Add(key, value)
		}

		response := httptest.NewRecorder()
		KrakenTestServer(response, request)

		got := response.Body.String()
		want := "{}"
		if got != want {
			t.Errorf("got %+v, want %q", got, want)
		}
	})

}
