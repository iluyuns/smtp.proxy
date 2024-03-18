package client

import (
	"errors"
	"regexp"

	"gopkg.in/gomail.v2"
)

type Client struct {
	Identity string
	Username string
	Password string
	Host     string
	Port     int
}

// CheckEmail 检查邮箱格式
func (i *Client) CheckEmail(email string) error {
	// 匹配所有子域名或顶级域名
	domainRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z.]{2,}`
	domainMatch, domainErr := regexp.MatchString(domainRegex, email)
	if domainErr != nil {
		return domainErr
	}
	if !domainMatch {
		return errors.New("email format error: domain not matched")
	}

	return nil
}

func (i *Client) SendMailWithTLS(from string, to []string, subject, body string, filename ...string) error {
	// 检查 from 是否符合邮箱格式
	if err := i.CheckEmail(from); err != nil {
		return err
	}
	for _, v := range to {
		// 检查 to 是否符合邮箱格式
		if err := i.CheckEmail(v); err != nil {
			return err
		}
	}
	d := gomail.NewDialer(i.Host, i.Port, i.Username, i.Password)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	// m.SetAddressHeader("Cc", "i@example.com", "i") 抄送
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if len(filename) != 0 {
		for _, v := range filename {
			m.Attach(v)
		}
	}
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
