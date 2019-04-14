package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

const TokenTimeOut = 20 * time.Minute

var (
	tokenTimestamp time.Time
	cachedToken    string
	locker         sync.Mutex
)

type (
	Config struct {
		BaseURL      string
		APIKey       string
		SecretKey    string
		DisableCache bool
	}

	BulkSMSProvider interface {
		GetCredit() (int, error)
	}

	TokenProvider interface {
		Get() (string, error)
	}

	Token struct {
		baseURL      string
		apiKey       string
		secretKey    string
		disableCache bool
	}

	bulkSMS struct {
		BaseURL string
		Token   TokenProvider
	}

	creditResult struct {
		Credit       int    `json:"Credit"`
		IsSuccessful bool   `json:"IsSuccessful"`
		Message      string `json:"Message"`
	}

	tokenResult struct {
		TokenKey     string `json:"TokenKey"`
		IsSuccessful bool   `json:"IsSuccessful"`
		Message      string `json:"Message"`
	}
)

func NewBulkSMSClient(token TokenProvider, url string) BulkSMSProvider {
	return &bulkSMS{
		BaseURL: url,
		Token:   token,
	}
}

func NewToken(config *Config) TokenProvider {
	return &Token{
		baseURL:      config.BaseURL,
		apiKey:       config.APIKey,
		secretKey:    config.SecretKey,
		disableCache: config.DisableCache,
	}
}

func (t *Token) Get() (string, error) {
	if !t.disableCache && (time.Now().Sub(tokenTimestamp) < TokenTimeOut) {
		return cachedToken, nil
	}

	locker.Lock()
	defer locker.Unlock()

	url := t.baseURL + "/token"
	data := struct {
		UserApiKey string `json:"UserApiKey"`
		SecretKey  string `json:"SecretKey"`
	}{UserApiKey: t.apiKey, SecretKey: t.secretKey}
	b, _ := json.Marshal(&data)

	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
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

func (b *bulkSMS) GetCredit() (int, error) {
	url := b.BaseURL + "/credit"
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	token, _ := b.Token.Get()
	r.Header.Add("x-sms-ir-secure-token", token)
	r.Header.Add("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(r)
	data := creditResult{}
	_ = json.NewDecoder(resp.Body).Decode(&data)
	defer resp.Body.Close()

	if data.IsSuccessful {
		return data.Credit, nil
	}
	return 0, errors.New("invalid token")
}
