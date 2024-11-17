package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	addr := os.Getenv("SMTP_ADDR")
	host := os.Getenv("SMTP_HOST")

	if from == "" || password == "" || addr == "" || host == "" {
		log.Fatal("Required environment variables are missing")
	}
	user := "contact.pecet.it@gmail.com"
	t := fmt.Sprintf("To: %s\r\n", to)
	s := fmt.Sprintf("Subject: %s\r\n", subject)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := []byte(t +

		s +

		mime +

		body + "\r\n")

	auth := smtp.PlainAuth(
		"",
		user,
		password,
		host)

	err := smtp.SendMail(
		addr,
		auth,
		from,
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		log.Println("Unable to send an email")
		return err
	}
	log.Printf("@email with subject:%s has been sent to: %s", subject, to)
	return nil
}
