package smtp

import "testing"

const (
	tos        = "ulric@b.com;rain@c.com"
	subject    = "test"
	body       = "test"
	SkipVerify = true
)

func Test_SendMailByTLS(t *testing.T) {
	address := "smtp.qq.com:465"
	from := "notify@qq.com"
	password := "password"
	tls := true
	anonymous := false
	s := NewSMTP(address, from, password, tls, anonymous, SkipVerify)
	t.Log(s.SendMail(from, tos, subject, body))
}

func Test_SendMail(t *testing.T) {
	address := "smtp.163.com:25"
	from := "notify@163.com"
	password := "password"
	s := New(address, from, password)
	t.Log(s.SendMail(from, tos, subject, body))
}

func Test_SendMailAnonymous(t *testing.T) {
	address := "localhost:25"
	from := "notify@localhost"
	tls := false
	anonymous := true
	s := NewSMTP(address, "", "", tls, anonymous, SkipVerify)
	t.Log(s.SendMail(from, tos, subject, body))
}
