package http

import (
	"crypto/tls"
	"net/http"

	"github.com/certifi/gocertifi"
)

func init() {
	certPool, err := gocertifi.CACerts()
	if err != nil {
		panic(err)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{RootCAs: certPool}
}
