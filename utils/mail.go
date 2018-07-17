package utils

import (
	"github.com/xblymmx/huzhi123/config"
	"fmt"
	"net/smtp"
	"crypto/tls"
	"net"
)

func SendMail(to, subject, content string) error {
	host := config.ServerConfig.MailHost
	port := config.ServerConfig.MailPort
	user := config.ServerConfig.MailUser
	pwd := config.ServerConfig.MailPassword
	from := config.ServerConfig.MailFrom

	headers := map[string]string{
		"From":         from,
		"To":           to,
		"Subject":      subject,
		"Content-Type": "text/html charset=utf-8",
	}

	msg := ""
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += content

	auth := smtp.PlainAuth("", user, pwd, host)

	err := sendMailThroughTLS(
		fmt.Sprintf("%s:%s", host, port),
		auth,
		from,
		[]string{to},
		msg,
	)

	return err
}

func sendMailThroughTLS(addr string, auth smtp.Auth, from string, to []string, msg string) (err error) {
	client, err := newSMTPClient(addr)
	if err != nil {
		return err
	}
	defer client.Close()

	if auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(auth); err != nil {
				return err
			}
		}
	}

	err = client.Mail(from)
	if err != nil {
		return err
	}

	for _, v := range to {
		if err = client.Rcpt(v); err != nil {
			fmt.Println(err)
			return err
		}
	}

	data, err := client.Data()
	if err != nil {
		return err
	}

	_, err = data.Write([]byte(msg))
	if err != nil {
		return err
	}

	err = data.Close()
	if err != nil {
		return err
	}

	return client.Quit()
}

func newSMTPClient(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tls", addr, nil)
	if err != nil {
		return nil, err
	}

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}
