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

func TestGetBalance(t *testing.T) {

	t.Run("returns balance", func(t *testing.T) {
		values := url.Values{}
		secret, _ := base64.StdEncoding.DecodeString(os.Getenv("API_KRAKEN_SECRET"))
		values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))
		headers := map[string]string{
			"API-Key":  os.Getenv("API_KRAKEN_KEY"),
			"API-Sign": utils.Signature("/0/private/Balance", values, secret),
		}
		request, _ := http.NewRequest(http.MethodPost, "https://api.kraken.com/0/private/Balance", strings.NewReader(values.Encode()))
		request.Header.Add("User-Agent", "GoKrakenBot/1.0")
		for key, value := range headers {
			request.Header.Add(key, value)
		}
		response := httptest.NewRecorder()
		KrakenTestServer(response, request)
		AssertResponseBody(t, response.Body.String(), "{20}")
	})

}
