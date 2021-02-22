package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"
)

// TestSignature is testing signature
func TestSignature(t *testing.T) {

	values := url.Values{}

	urlPath := fmt.Sprintf("/0/private/Balance")
	secretKey, _ := base64.StdEncoding.DecodeString(os.Getenv("API_KRAKEN_SECRET"))

	// set nonce
	nonce := fmt.Sprintf("%d", time.Now().UnixNano())
	values.Set("nonce", nonce)

	secret := []byte(secretKey)
	shaSum := getSha256Test([]byte(values.Get("nonce") + values.Encode()))
	macSum := getHMacSha512Test(append([]byte(urlPath), shaSum...), secret)
	t.Logf("nonce: %s", nonce)
	t.Logf("signature: %s", base64.StdEncoding.EncodeToString(macSum))
}

// getSha256 creates a sha256 hash for given []byte
func getSha256Test(input []byte) []byte {
	sha := sha256.New()
	sha.Write(input)
	return sha.Sum(nil)
}

// getHMacSha512 creates a hmac hash with sha512
func getHMacSha512Test(message, secret []byte) []byte {
	mac := hmac.New(sha512.New, secret)
	mac.Write(message)
	return mac.Sum(nil)
}
