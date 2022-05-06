package lib

import (
	"crypto/tls"
	"github.com/prometheus/common/expfmt"
	"net/http"
)

type PromPull struct {
	URI  string
	USER string
	PASS string
}

func (pp *PromPull) Pull(metric string) float64 {

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
	// body, _ := ioutil.ReadAll(resp.Body)
	var parser expfmt.TextParser
	prom, _ := parser.TextToMetricFamilies(resp.Body)
	found := prom[metric]
	if found != nil {
		return prom[metric].Metric[0].GetGauge().GetValue()
	}
	return -1

}
