<img src="http://crypdex.io/img/full-logo.svg" width=300 style="margin-bottom:20px;"/>

# A Go package for [CoinSwitch](https://developer.coinswitch.co/reference)

Golang wrapper for the CoinSwitch instant coin exchange. Implements API v2.


```bash
go get github.com/crypdex/go-coinswitch
```

## Usage

```go
import "github.com/crypdex/go-coinswitch"

// Initialize the client
apikey := "gogetone"
client := coinswitch.New(apikey)

// Make a request (catch your error)
rate, _ := client.GetRate("pivx", "tusd")

fmt.Println(rate)

// {
//   "rate": 0.75239778,
//   "minerFee": 0.7,
//   "limitMinDepositCoin": 40,
//   "limitMaxDepositCoin": 7850,
//   "limitMinDestinationCoin": 30.09591127,
//   "limitMaxDestinationCoin": 5906.32258769
// }
```

This package is currently used in production, but is not yet locked down by version and is subject to change.

## Implemention Status (v2)

- [x] Dynamic API
- [ ]  Fixed Rate API
    
## Contributing

* Fork the repo
* Add desired functionality
* Issue a Pull Request

Better testing with mocks needs to be written.

## 🐞 Bugs

If you find a problem with the lib, please submit an issue.