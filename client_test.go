package a1exchange

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	c := New()

	for i := 0; i < 10; i++ {
		res, err := c.Ticker()
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%+v\n", res)
		time.Sleep(time.Second)
	}
}

func TestDepth(t *testing.T) {
	c := New()

	for i := 0; i < 10; i++ {
		res, err := c.Depth(VEOBTC)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%+v\n", res)
		time.Sleep(time.Second)
	}
}

func TestTrades(t *testing.T) {
	c := New()

	for i := 0; i < 10; i++ {
		res, err := c.Trades(VEOBTC)
		if err != nil {
			t.Error(err)
		}

		for _, v := range res.Trades {
			fmt.Printf("%+v\n", v)
		}
		time.Sleep(time.Second)
	}
}
