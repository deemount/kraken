package utils

import (
	"fmt"
	"strings"
)

// Error represents a Kraken API error
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("kraken error: %s", strings.ToLower(e.Message))
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
