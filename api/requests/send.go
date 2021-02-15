package requests

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// HTTPClient is a struct
type HTTPClient struct {
	client http.Client
}

// Send is a method
func (c *HTTPClient) Send(url string, params url.Values, headers map[string]string) ([]byte, error) {

	var err error

	log.Printf("send request uri: \n%v\n%v", url, strings.NewReader(params.Encode()))

	// initialize request
	req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
	}

	req.Header.Add("User-Agent", "GoKrakenBot/1.0")

	// add headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	log.Printf("%+v", req)

	// execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return []byte("0"), err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// check response status
	// if resp.StatusCode != http.StatusOK {
	// 	switch resp.StatusCode {
	// 	case http.StatusUnauthorized:
	// 		return []byte("0"), utils.ErrUnauthorized
	// 	case http.StatusNotFound:
	// 		return []byte("0"), utils.ErrNotFound
	// 	default:
	// 		return []byte("0"), utils.ErrDefaultResponse // resp.Status
	// 	}
	// }

	// check response mime type
	// mimeType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	// if err != nil {
	// 	log.Fatal("Error on parsing mime type. ", err.Error())
	// }
	// if mimeType != "application/json" {
	// 	log.Fatalf("Error on mime type. Response Content-Type is '%s', but should be 'application/json'.", mimeType)
	// }

	// return response body
	return body, err

}
