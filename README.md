# go-a1exchange

A1 Exchange API v0.2

## Description

go-a1exchange is a go client library for [A1 Exchange API v0.2](https://app.swaggerhub.com/apis/A1-Exchange/a1exchange/0.2.0).

## Installation

```
$ go get -u github.com/go-numb/go-a1exchange
```

## Usage
``` 
package main

import (
 "fmt"
 "github.com/go-numb/go-a1exchange"
)


func main() {
	c := New()

    res, err := c.Depth(VEOBTC)
    if err != nil {
        t.Error(err)
    }

    fmt.Printf("%+v\n", res)
}
```

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-a1exchange/blob/master/LICENSE)