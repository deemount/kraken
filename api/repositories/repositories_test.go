package repositories

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func KrakenTestServer(w http.ResponseWriter, r *http.Request) {
	method := strings.TrimPrefix(r.URL.Path, "/0/private/")

	if method == "Ledger" {
		fmt.Fprint(w, "{30}")
		return
	}

	if method == "Balance" {
		fmt.Fprint(w, "{20}")
		return
	}

	if method == "TradeBalance" {
		fmt.Fprint(w, "{10}")
		return
	}

}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
