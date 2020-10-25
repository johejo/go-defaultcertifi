package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/johejo/go-defaultcertifi/http"
)

func main() {
	resp, err := http.Get("https://api.ssllabs.com/api/v3/getRootCertsRaw")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
