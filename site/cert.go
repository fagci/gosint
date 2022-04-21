package site

import (
	"crypto/tls"
	"crypto/x509"
	"net/url"
	"regexp"
	"sort"
)

var protoRegexp = regexp.MustCompile(`^(\w)://`)

func DNSNames(addr string) (dnsNames []string) {
	addr = "https://" + protoRegexp.ReplaceAllString(addr, "")

	if uri, err := url.Parse(addr); err == nil {
		certs, _ := Certs(uri)
		for _, cert := range certs {
			dnsNames = append(dnsNames, cert.DNSNames...)
		}
	}

    sort.Strings(dnsNames)

	return
}

func Certs(uri *url.URL) ([]*x509.Certificate, error) {
	if uri.Port() == "" {
		uri.Host += ":443"
	}
	conn, err := tls.Dial("tcp", uri.Host, nil)
	if err != nil {
		return nil, err
	}

	return conn.ConnectionState().PeerCertificates, nil
}
