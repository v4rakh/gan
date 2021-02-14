package mail

import (
	"github.com/v4rakh/gan/internal/gan/constant"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"os"
	"strconv"
	"time"
)

type Service struct {
	smtpServer *mail.SMTPServer
}

func NewService() *Service {
	if os.Getenv(constant.EnvMailEnabled) == "false" {
		return &Service{
			smtpServer: nil,
		}
	}

	port, err := strconv.Atoi(os.Getenv(constant.EnvMailPort))
	if err != nil {
		log.Fatalf("Mail port is not valid. Reason: %s\n", err.Error())
	}

	server := mail.NewSMTPClient()
	server.Host = os.Getenv(constant.EnvMailHost)
	server.Port = port
	server.Username = os.Getenv(constant.EnvMailAuthUser)
	server.Password = os.Getenv(constant.EnvMailAuthPassword)

	if os.Getenv(constant.EnvMailEncryption) == "NONE" {
		server.Encryption = mail.EncryptionNone
	} else if os.Getenv(constant.EnvMailEncryption) == "TLS" {
		server.Encryption = mail.EncryptionTLS
	} else {
		server.Encryption = mail.EncryptionSSL
	}

	if os.Getenv(constant.EnvMailEncryption) == "LOGIN" {
		server.Authentication = mail.AuthLogin
	} else if os.Getenv(constant.EnvMailEncryption) == "CRAM_MD5" {
		server.Authentication = mail.AuthCRAMMD5
	} else {
		server.Authentication = mail.AuthPlain
	}

	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	return &Service{
		smtpServer: server,
	}
}

func (s *Service) Send(address string, subject string, body string) {
	if os.Getenv(constant.EnvMailEnabled) == "false" {
		return
	}

	email := mail.NewMSG()
	email.SetFrom(os.Getenv(constant.EnvMailFrom)).AddTo(address).SetSubject(subject).SetBody(mail.TextPlain, body)

	smtpClient, err := s.smtpServer.Connect()
	if err != nil {
		log.Fatalf("Could not connect to mail server '%s'. Reason: %s\n", s.smtpServer.Host, err.Error())
	}

	err = email.Send(smtpClient)
	if err != nil {
		log.Printf("Could not send mail to '%s'. Reason: %s\n", address, err.Error())
	} else {
		log.Printf("Mail to '%s' with subject '%s' sent\n", address, subject)
	}
}
