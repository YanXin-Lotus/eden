package services

import (
	"crypto/md5"
	"encoding/hex"
	"net/smtp"
)

func Md5(origin string) string {
	hasher := md5.New()
	hasher.Write([]byte(origin))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SendMail(to string, sub string, body string) (err error) {
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
	msg := []byte("To: " + to + "\n" + "Subject: " + sub + "\r\n" + "\r\n" + body + "\r\n")
	err = smtp.SendMail("mail.example.com:25", auth, "sender@example.com", to, msg)
	return err
}
