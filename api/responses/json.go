// Copyright 2021. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

// Package responses ...
package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON encodes response data
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

}
