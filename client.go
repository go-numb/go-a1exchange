package a1exchange

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	// curl -X GET "https://api.a1.exchange/v1/tickers" -H "accept: application/json"
	ENDPOINT = "https://api.a1.exchange"
	VERSION  = "v1"

	VEOBTC = "VEOBTC"
	VEOETH = "VEOETH"
)

type Client struct {
	BaseURL string

	HTTPClient *http.Client
}

func New() *Client {
	u, err := url.Parse(ENDPOINT)
	if err != nil {
		log.Fatal(err)
	}

	u.Path = VERSION

	client := http.DefaultClient
	tr := &http.Transport{
		MaxIdleConnsPerHost: 24,
		TLSHandshakeTimeout: 0 * time.Second,
	}
	client = &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	return &Client{
		BaseURL:    u.String(),
		HTTPClient: client,
	}
}
