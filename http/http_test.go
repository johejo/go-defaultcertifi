package http_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	_ "github.com/johejo/go-defaultcertifi/http"
)

func Test_init(t *testing.T) {
	resp, err := http.Get("https://api.ssllabs.com/api/v3/getRootCertsRaw")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(b))
}
