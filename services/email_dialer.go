package services

import (
	"crypto/tls"
	"github.com/codersgarage/smart-cashier/config"
	"github.com/go-gomail/gomail"
)

func EmailDialer() *gomail.Dialer {
	cfg := config.EmailService()
	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUsername, cfg.SMTPPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d
}
