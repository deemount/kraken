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

func KrakenBalanceServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

func TestGetBalance(t *testing.T) {
	t.Run("returns balance", func(t *testing.T) {

		values := url.Values{}
		secret, _ := base64.StdEncoding.DecodeString(os.Getenv("API_KRAKEN_SECRET"))
		values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

		headers := map[string]string{
			"API-Key":  os.Getenv("API_KRAKEN_KEY"),
			"API-Sign": utils.Signature("/0/private/Balance", values, secret),
		}

		t.Logf("%v", headers)

		request, _ := http.NewRequest(http.MethodPost, "https://api.kraken.com/0/private/Balance", strings.NewReader(values.Encode()))
		request.Header.Add("User-Agent", "GoKrakenBot/1.0")
		for key, value := range headers {
			request.Header.Add(key, value)
		}

		response := httptest.NewRecorder()

		KrakenBalanceServer(response, request)

		// var client http.Client
		// response, err := client.Do(request)
		// if err != nil {
		// 	t.Errorf("%v", err)
		// }
		// defer response.Body.Close()

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %+v, want %q", got, want)
		}
	})
}
