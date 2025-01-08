package logger

import (
	"context"
	"fmt"
	"net/smtp"
)

type Sender interface {
	Send(ctx context.Context, l *Logger) error
}

func (e Email) Send(ctx context.Context, l *Logger) error {
	subject := fmt.Sprintf("Subject: %s\r\n", e.Subject)
	body := fmt.Sprintf("%s\r\n", e.BodyTemplate)
	msg := subject + "\r\n" + body

	addr := fmt.Sprintf("%s:%d", e.SMTPHost, e.SMTPPort)

	done := make(chan error, 1)
	go func() {
		auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPHost)
		err := smtp.SendMail(addr, auth, e.FromAddress, e.ToAddresses, []byte(msg))
		done <- err
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("email sending canceled: %w", ctx.Err())
	case err := <-done:
		if err != nil {
			return err
		}

	}

	return nil
}
