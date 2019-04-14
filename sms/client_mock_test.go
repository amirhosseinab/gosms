package sms_test

import "github.com/amirhosseinab/go-sms-ir/sms"

type fakeToken struct {
	token string
}

func createFakeToken(token string) sms.TokenProvider {
	return &fakeToken{token: token}
}
func (t *fakeToken) Get() (string, error) {
	return t.token, nil
}
