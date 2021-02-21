package repositories

import (
	"fmt"
	"net/http"
)

func KrakenTestServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}
