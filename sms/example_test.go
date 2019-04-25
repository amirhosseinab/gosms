package sms_test

import (
	"fmt"
	"github.com/amirhosseinab/go-sms-ir/sms"
	"log"
)

func ExampleBulkSMS_GetCredit() {
	tp := sms.NewToken(sms.Config{
		APIKey:    "YOUR_API_KEY",
		SecretKey: "YOUR_SECRET_KEY",
	})

	client := sms.NewBulkSMSClient(tp, sms.DefaultBulkURL)

	credit, err := client.GetCredit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(credit)
}
