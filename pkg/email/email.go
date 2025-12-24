package email

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Interface interface {
	SendRestPasswordEmail(toEmail, token string) error
}

type Email struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func Init() Interface {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	return &Email{
		host:     os.Getenv("SMTP_HOST"),
		port:     port,
		username: os.Getenv("SMTP_USERNAME"),
		password: os.Getenv("SMTP_PASSWORD"),
		from:     os.Getenv("SMTP_FROM"),
	}
}

func (e *Email) SendRestPasswordEmail(toEmail, token string) error {
	resetURL := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("APP_URL"), token)

	// message
	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Reset Password - My App")

	// HTML Body
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Reset Password</h2>
			<p>Klik link di bawah untuk reset password:</p>
			<p><a href="%s">Reset Password</a></p>
			<p>Link ini expired dalam 1 jam.</p>
		</body>
		</html>
	`, resetURL)

	m.SetBody("text/html", body)

	// Send email
	d := gomail.NewDialer(e.host, e.port, e.username, e.password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %s", err)
	}

	return nil
}
