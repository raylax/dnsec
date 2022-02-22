package client

import (
	"encoding/base64"
	"github.com/miekg/dns"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Client interface {
	Query(*dns.Msg) (*dns.Msg, error)
}

type dotClient struct {
	endpoint string
	cli *dns.Client
}

func NewDoT(endpoint string) *dotClient {
	c := &dotClient{
		endpoint: endpoint,
		cli: &dns.Client{
			Net: "tcp-tls",
		},
	}
	return c
}

func (c *dotClient) Query(q *dns.Msg) (r *dns.Msg, err error) {
	r, _, err = c.cli.Exchange(q, c.endpoint)
	return
}

type dohClient struct {
	endpoint string
}

func NewDoH(endpoint string) *dohClient {
	return &dohClient{endpoint: endpoint}
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

func (c *dohClient) Query(q *dns.Msg) (r *dns.Msg, err error) {
	bytes, err := q.Pack()
	if err != nil {
		return
	}
	data := base64.RawURLEncoding.EncodeToString(bytes)
	req, err := http.NewRequest("GET", c.endpoint + "?dns=" + data, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/dns-message")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	r = &dns.Msg{}
	r.SetReply(q)
	err = r.Unpack(result)
	return
}
