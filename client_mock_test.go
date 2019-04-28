package gosms_test

import "github.com/amirhosseinab/gosms"

type fakeToken struct {
	token string
}

func createFakeToken(token string) gosms.TokenProvider {
	return &fakeToken{token: token}
}
func (t *fakeToken) Get() (string, error) {
	return t.token, nil
}
