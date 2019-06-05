package a1exchange

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Depth struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

// https://api.a1.exchange/v1/markets/VEOBTC/depth
func (p *Client) Depth(code string) (*Depth, error) {
	u, _ := url.Parse(p.BaseURL)
	u.Path += fmt.Sprintf("/markets/%s/depth", code)
	res, err := p.HTTPClient.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	d := new(Depth)
	json.NewDecoder(res.Body).Decode(d)

	return d, nil
}
