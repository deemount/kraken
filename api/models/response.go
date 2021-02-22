// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package models ...
package models

// Response ...
type Response struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}
