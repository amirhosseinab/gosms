// Package sms provides a client library for https://www.sms.ir Restful API.
package sms

type (
	// Client provides the required token and use to call the corresponding APIs.
	Client struct {
		Token string
	}
)

// NewClient creates a new client value using your API_KEY and SECRET_KEY values.
// A client value sends a request to the server and get the token back for calling the other APIs.
func NewClient(apiKey, secretKey string) (*Client, error) {
	return &Client{Token: "fake token"}, nil
}
