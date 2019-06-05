package a1exchange

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type T struct {
	Code   string
	Trades []Trade `json:"trades"`
}
type Trade struct {
	ID     int     `json:"id"`
	Price  float64 `json:"price,string"`
	Time   float64 `json:"time"`
	Amount float64 `json:"amount,string"`
	Type   string  `json:"type"`
}

func (p *Client) Trades(code string) (*T, error) {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Println("exec time: ", end.Sub(start))
	}()
	u, _ := url.Parse(p.BaseURL)
	u.Path += fmt.Sprintf("/markets/%s/trades", code)
	res, err := p.HTTPClient.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Printf("%+v\n", string(body))

	t := new(T)
	json.NewDecoder(res.Body).Decode(t)
	t.Code = code

	return t, nil
}
