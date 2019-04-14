package sms_test

import (
	"github.com/amirhosseinab/go-sms-ir/sms"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := sms.NewClient("", "")
	if err != nil {
		t.Fatal(err)
	}
	if client.Token == "" {
		t.Fatal("token is null")
	}
}
