package site

import (
	"fmt"
	"testing"
)

func TestCertDomains(t *testing.T) {
	for _, domain := range DNSNames("ya.ru") {
		fmt.Println(domain)
	}
}
