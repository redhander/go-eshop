package mailer

import (
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

const (
	smtpGmailHost   = "smtp.gmail.com"
	smtpGmailPort   = 587
	smtpMailHogHost = "localhost"
	smtpMailHogPort = 1025
)

type EmailSender interface {
	Send(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type emailSender struct {
	dialer   *gomail.Dialer
	from     string
	host     string
	port     int
	username string
}

func NewEmailSender(host string, port int, username, password, from, env string) EmailSender {
	// If we have proper SMTP settings, use them
	if host != "" && port != 0 && username != "" && password != "" {
		dialer := gomail.NewDialer(host, port, username, password)

		// Special handling for SSL ports like 465
		if port == 465 {
			dialer.SSL = true
		}

		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		// Log configuration in development mode for debugging
		if env == "development" {
			log.Printf("SMTP Configuration: Host=%s, Port=%d, Username=%s, SSL=%t", host, port, username, dialer.SSL)
		}

		return &emailSender{
			dialer:   dialer,
			from:     from,
			host:     host,
			port:     port,
			username: username,
		}
	}

	// Fallback to MailHog for development if no credentials provided
	if env == "development" {
		log.Printf("Using MailHog fallback: Host=%s, Port=%d", smtpMailHogHost, smtpMailHogPort)
		return &emailSender{
			dialer: &gomail.Dialer{
				Host: smtpMailHogHost,
				Port: smtpMailHogPort,
			},
			from: "email.admin@eshop.com",
		}
	}

	// Fallback to Gmail settings
	log.Printf("Using Gmail fallback: Host=%s, Port=%d, Username=%s", smtpGmailHost, smtpGmailPort, username)
	dialer := gomail.NewDialer(smtpGmailHost, smtpGmailPort, username, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &emailSender{
		dialer:   dialer,
		from:     username,
		host:     smtpGmailHost,
		port:     smtpGmailPort,
		username: username,
	}
}

func (g *emailSender) Send(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	mail := gomail.NewMessage()

	from := g.from
	if from == "" {
		from = "email.admin@eshop.com"
	}

	mail.SetHeaders(
		map[string][]string{
			"From":    {from},
			"To":      to,
			"Subject": {subject},
		},
	)
	mail.SetHeader("Cc", cc...)
	mail.SetHeader("Bcc", bcc...)
	mail.SetBody("text/html", content)

	if err := g.dialer.DialAndSend(mail); err != nil {
		// Add more context to the error message
		return fmt.Errorf("failed to send email using SMTP server %s:%d with username %s: %w",
			g.host, g.port, g.username, err)
	}
	return nil
}
