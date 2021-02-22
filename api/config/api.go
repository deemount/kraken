// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package config ...
package config

// API ...
type API struct {
	Version     int
	Host        string
	Path        string
	Port        string
	StrictSlash bool
	LimitClient int
	Destination string
	Msg         Message
}

// Message ...
type Message struct {
	Status int
	Notice string
	Out    string
}
