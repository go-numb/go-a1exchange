package a1exchange

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Data map[string]OHLC

type OHLC struct {
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	Last   float64 `json:"last"`
	High   float64 `json:"high"`
	Volume float64 `json:"volume"`
}

type Ticker struct {
	Code string
	OHLC
}

func (p *Client) Ticker() ([]Ticker, error) {
	u, _ := url.Parse(p.BaseURL)
	u.Path += "/tickers"
	res, err := p.HTTPClient.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Printf("%+v\n", string(body))

	var t Data
	json.NewDecoder(res.Body).Decode(&t)

	// fmt.Printf("%+v\n", t)

	// var tickers Tickers

	// 使いやすい形に変更
	tickers := make([]Ticker, len(t))
	var count int
	for key, val := range t {
		// fmt.Printf("%s :::::::: %+v\n", key, val)
		tickers[count] = Ticker{
			key,
			val,
		}
		count++
	}

	return tickers, nil
}

// tcp用
func discard(res *http.Response) {
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
}
