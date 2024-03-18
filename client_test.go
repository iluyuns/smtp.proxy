package client_test

import (
	"testing"

	client "github.com/iluyuns/smtp.proxy"
)

// TestCheckEmail
func TestCheckEmail(t *testing.T) {
	client := &client.Client{}
	tests := []string{
		"49@qq.com",
		"ss@e.sss.com",
		"ee@e.e.e.e.app",
		"e@e.com.ssss",
	}
	for _, v := range tests {
		err := client.CheckEmail(v)
		if err != nil {
			t.Error(err)
		}
	}
}
