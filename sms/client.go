package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

// TokenTimeOut indicates refresh time of the token for accessing APIs.
const TokenTimeOut = 20 * time.Minute

// DefaultBulkURL is used to send requests to SMS provider by default.
const DefaultBulkURL = "https://restfulsms.com/api"

var (
	tokenTimestamp time.Time
	cachedToken    string
	locker         sync.Mutex
)

type (
	// Config holds the data that is required for constructing Token.
	Config struct {
		BaseURL      string
		APIKey       string
		SecretKey    string
		DisableCache bool
	}

	// BulkSMSProvider exposes the methods of bulk SMS system.
	//
	// GetCredit fetch the amount of the SMS count that remains on the account.
	// It uses the token that provides by the Token.Get() method.
	BulkSMSProvider interface {
		GetCredit() (int, error)
	}

	// TokenProvider is used to fetch the token from the server.
	TokenProvider interface {
		Get() (string, error)
	}

	// Token handles the requests for providing token
	Token struct {
		Config Config
	}

	BulkSMS struct {
		BaseURL string
		Token   TokenProvider
	}

	creditResult struct {
		Credit       float32 `json:"Credit"`
		IsSuccessful bool    `json:"IsSuccessful"`
		Message      string  `json:"Message"`
	}

	tokenResult struct {
		TokenKey     string `json:"TokenKey"`
		IsSuccessful bool   `json:"IsSuccessful"`
		Message      string `json:"Message"`
	}
)

// NewBulkSMSClient create a value that handle all requests for bulk SMS system.
func NewBulkSMSClient(token TokenProvider, url string) BulkSMSProvider {
	return &BulkSMS{
		BaseURL: url,
		Token:   token,
	}
}

// NewToken provides a value for fetching token from the server.
func NewToken(config Config) TokenProvider {
	url := config.BaseURL
	if url == "" {
		url = DefaultBulkURL
	}

	token := &Token{Config: config}
	token.Config.BaseURL = url
	return token
}

// Get method fetches token from the server.
// It is thread safe and handles the caching mechanism by default to prevent unnecessary requests.
func (t *Token) Get() (string, error) {
	if !t.Config.DisableCache && (time.Now().Sub(tokenTimestamp) < TokenTimeOut) {
		return cachedToken, nil
	}

	locker.Lock()
	defer locker.Unlock()

	url := t.Config.BaseURL + "/token"
	data := struct {
		UserApiKey string `json:"UserApiKey"`
		SecretKey  string `json:"SecretKey"`
	}{UserApiKey: t.Config.APIKey, SecretKey: t.Config.SecretKey}
	b, _ := json.Marshal(&data)

	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	r.Header.Add("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(r)
	result := tokenResult{}
	_ = json.NewDecoder(resp.Body).Decode(&result)
	if result.IsSuccessful {
		cachedToken = result.TokenKey
		tokenTimestamp = time.Now()

		return result.TokenKey, nil
	}
	return "", errors.New("invalid API key or secret key")
}

// GetCredit fetch the amount of the SMS count that remains on the account.
// It uses the token that provides by the Token.Get() method.
func (b *BulkSMS) GetCredit() (int, error) {
	url := b.BaseURL + "/credit"
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	token, _ := b.Token.Get()
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("x-sms-ir-secure-token", token)
	resp, _ := http.DefaultClient.Do(r)
	data := creditResult{}
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if data.IsSuccessful {
		return int(data.Credit), nil
	}
	return 0, errors.New("invalid token")
}
