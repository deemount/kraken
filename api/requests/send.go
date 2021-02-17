package requests

import (
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strings"

	"github.com/deemount/kraken/api/utils"
)

// HTTPClient is a struct
type HTTPClient struct {
	client http.Client
}

// Send is a method
func (c *HTTPClient) Send(url string, params url.Values, headers map[string]string) ([]byte, error) {

	var err error

	// initialize request
	req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	if err != nil {
		return []byte("0"), utils.ErrCreateObject
	}

	// add useragent
	req.Header.Add("User-Agent", "GoKrakenBot/1.0")

	// add headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return []byte("0"), err
	}
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return []byte("0"), utils.ErrUnauthorized
		case http.StatusNotFound:
			return []byte("0"), utils.ErrNotFound
		default:
			return []byte("0"), utils.ErrDefaultResponse
		}
	}

	// check response mime type
	mimeType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return []byte("0"), utils.ErrParseMimeType
	}
	if mimeType != "application/json" {
		return []byte("0"), utils.ErrRespMimeType
	}

	// return response body
	return ioutil.ReadAll(resp.Body)

}
