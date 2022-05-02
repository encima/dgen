package lib

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PromPull struct {
	URI  string
	USER string
	PASS string
}

func (pp *PromPull) Pull() {

	req, err := http.NewRequest("GET", pp.URI, nil)
	req.SetBasicAuth(pp.USER, pp.PASS)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{Transport: tr}
	resp, err := cli.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
