package sms_test

import (
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

	log.Printf("Your credit is: %d", credit)
}

func ExampleBulkSMS_SendVerificationCode() {
	token := sms.NewToken(sms.Config{
		APIKey:    "YOUR_API_KEY",
		SecretKey: "YOUR_SECRET_KEY",
	})

	client := sms.NewBulkSMSClient(token, sms.DefaultBulkURL)

	vId, err := client.SendVerificationCode("09121234567", "123456")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Verification Id is %s", vId)
}

func ExampleBulkSMS_SendByTemplate() {
	token := sms.NewToken(sms.Config{
		APIKey:    "YOUR_API_KEY",
		SecretKey: "YOUR_SECRET_KEY",
	})

	client := sms.NewBulkSMSClient(token, sms.DefaultBulkURL)

	vId, err := client.SendByTemplate("09121234567", 1234, map[string]string{"VerificationCode": "9999"})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Verification Id is %s", vId)
}
