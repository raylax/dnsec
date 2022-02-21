package client

import (
	"encoding/base64"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Client struct {
	Endpoint string
}

var defaultTransport http.RoundTripper = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   3 * time.Second,
		KeepAlive: 10 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          10,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   1 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var client = &http.Client{
	Transport: defaultTransport,
}

func (cli Client) Query(q []byte) ([]byte, error) {
	data := base64.RawURLEncoding.EncodeToString(q)
	req, err := http.NewRequest("GET", cli.Endpoint + "?dns=" + data, nil)
	req.Header.Set("Content-Type", "application/dns-message")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
