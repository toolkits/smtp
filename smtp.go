package smtp

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

type Smtp struct {
	Address  string
	Username string
	Password string
}

func New(address, username, password string) *Smtp {
	return &Smtp{
		Address:  address,
		Username: username,
		Password: password,
	}
}

func (this *Smtp) SendMail(from, tos, subject, body string) error {
	if this.Address == "" {
		return fmt.Errorf("address is necessary")
	}

	hp := strings.Split(this.Address, ":")
	if len(hp) != 2 {
		return fmt.Errorf("address format error")
	}

	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = tos.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(subject)))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))

	auth := smtp.PlainAuth("", this.Username, this.Password, hp[0])
	return smtp.SendMail(this.Address, auth, from, strings.Split(tos, ";"), message)
}
