// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package config ...
package config

// DB is a struct
type DB struct {
	Driver        string
	User          string
	PW            string
	Port          string
	Host          string
	SSL           string
	Schema        string
	TblPrefix     string
	Name          string
	SingularTable bool
	LogMode       bool
}
