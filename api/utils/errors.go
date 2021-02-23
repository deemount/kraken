// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package utils ...
package utils

import (
	"strings"
)

// Error represents a Kraken API error
type Error struct {
	Message string
}

// Error is a method
func (e *Error) Error() string {
	return strings.ToLower(e.Message)
}

// APIErrors
var (
	ErrDefaultResponse      = &Error{"unknown error"}
	ErrUnauthorized         = &Error{"authentication required"}
	ErrCurrencyNotSupported = &Error{"currency does not supported"}
	ErrUserAgentInvalid     = &Error{"invalid user-agent"}
	ErrNotFound             = &Error{"not found"}
	ErrParseMimeType        = &Error{"parsing mime type"}
	ErrRespMimeType         = &Error{"response content-type must be 'application/json'"}
	ErrCreateObject         = &Error{"create request object not possible"}
)
