package coinswitch

import (
	"encoding/json"
	"fmt"
	"testing"
)

var apikey = ""

var client = New(apikey)

func TestCoinSwitch_Coins(t *testing.T) {
	result, err := client.GetCoins()
	if err != nil {
		t.Fatal(err)
	}

	print(result)
}

func TestCoinSwitch_Pairs(t *testing.T) {
	result, err := client.GetPairs("pivx", "")
	if err != nil {
		t.Fatal(err)
	}

	print(result)
}

func TestCoinSwitch_Rate(t *testing.T) {
	result, err := client.GetRate("pivx", "tusd")
	if err != nil {
		t.Fatal(err)
	}

	print(result)
}

func TestCoinSwitch_CreateOrder(t *testing.T) {
	request := &OrderRequest{
		DepositCoin:        "PIVX",
		DepositCoinAmount:  50,
		DestinationCoin:    "TUSD",
		DestinationAddress: Address{Address: "0x2D1Db67F39f959719c19b6CC3C83e19d6527eee4"},
	}
	result, err := client.CreateOrder(request)
	if err != nil {
		t.Fatal(err)
	}

	print(result)
}

// Sandbox Mode
// For ease of integration on client side, below are the list of orderId that can be sent to simulate which can be used to simulate the order lifecycle during API integration.
//
// 22222222-6c9e-4c53-9a6d-55e089aebd04 => no_deposit
// 33333333-6c9e-4c53-9a6d-55e089aebd04 => confirming
// 44444444-6c9e-4c53-9a6d-55e089aebd04 => exchanging
// 55555555-6c9e-4c53-9a6d-55e089aebd04 => sending
// 66666666-6c9e-4c53-9a6d-55e089aebd04 => complete
// 77777777-6c9e-4c53-9a6d-55e089aebd04 => failed
// 88888888-6c9e-4c53-9a6d-55e089aebd04 => refunded
// 11111111-6c9e-4c53-9a6d-55e089aebd04 => timeout

func TestCoinSwitch_GetOrderStatus(t *testing.T) {
	result, err := client.GetOrderStatus("22222222-6c9e-4c53-9a6d-55e089aebd04")
	if err != nil {
		t.Fatal(err)
	}

	print(result)
}

func TestCoinSwitch_GetOrders(t *testing.T) {
	result, err := client.GetOrders(0, 10)
	if err != nil {
		t.Fatal(err)
	}

	print(result)
}

func print(x interface{}) {
	j, _ := json.MarshalIndent(x, "", "  ")
	fmt.Println(string(j))
}
