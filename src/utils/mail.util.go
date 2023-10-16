package utils

import (
	"fmt"
	"go-food-delivery-api/src/models"
	"os"

	"gopkg.in/gomail.v2"
)

func SendSignUpSecretCode(to string, cc *string, account models.AccountCreatable) error {
	host := "smtp.gmail.com"
	port := 587
	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_APPLICATION_PASSWORD")
	from := os.Getenv("HOST_GMAIL_ADDRESS")
	sub := "NOTIFICATION ACCOUNT INFO EMAIL - WELCOME TO OUR SERVICE"
	body := fmt.Sprintf(`
		Welcome: %s to our service:
		Your Profile Information here:
			- Username: %s
			- Secret Code: %d
		Please keep the secret code carefully, you must use this secret code to publish an order or reset your password!
		Thank!
	`, *account.Username, *account.Username, *account.SecretCode)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", sub)
	m.SetBody("text/plain", body)
	if cc != nil {
		m.SetAddressHeader("Cc", *cc, *cc)
	}

	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error while send mail: " + err.Error())
		return err
	}
	return nil
}
