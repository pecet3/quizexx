package logger

import (
	"context"
	"errors"
	"fmt"
	"net/smtp"
)

type Sender interface {
	SendLogs(ctx context.Context, l *Logger) error
	SendAlert(ctx context.Context, l *Logger, content string) error
}

func (e Email) SendLogs(ctx context.Context, l *Logger) error {
	if len(l.cache) <= 0 {
		return errors.New("no logs in cache, canceled sending an email")
	}
	logs := ""
	for _, log := range l.cache {
		logs += log + "\n"
	}
	subject := fmt.Sprintf("Subject: %s\r\n", e.SubjectRaport)
	body := fmt.Sprintf("%s\r\n", logs)
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
		return errors.New("email sending canceled")
	case err := <-done:
		if err != nil {
			return err
		}

	}
	return nil
}
func (e Email) SendAlert(ctx context.Context, l *Logger, content string) error {
	subject := fmt.Sprintf("Subject: %s\r\n", e.SubjectAlert)
	body := fmt.Sprintf("%s\r\n", content)
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
