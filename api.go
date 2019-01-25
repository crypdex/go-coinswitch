package coinswitch

import (
	"errors"
	"strconv"
)

// Coins - This endpoint is used to retrieve all currencies supported by CoinSwitch
// https://developer.coinswitch.co/reference#apicoins
func (c *CoinSwitch) Coins() ([]Coin, error) {
	path := "/coins"
	result := new(CoinResult)
	if _, err := c.client.R().SetResult(result).Get(path); err != nil {
		return nil, err
	}
	return result.Data, nil
}

// Pairs - Get exchange pairs for a chosen depositCoin or destinationCoin
// https://developer.coinswitch.co/reference#apipairs
func (c *CoinSwitch) Pairs(depositCoin, destinationCoin string) ([]Pair, error) {
	path := "/pairs"
	result := new(PairsResult)

	_, err := c.client.R().
		SetResult(result).
		SetBody(map[string]string{
			"depositCoin":     depositCoin,
			"destinationCoin": destinationCoin,
		}).
		Post(path)

	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// Rate - Generate an exchange offer for for a coin pair
// https://developer.coinswitch.co/reference#apirate
//
// We are making an assumption here that if there is no data returned, but the request is successful
// that there is no trading pair for the requested deposit/destination
func (c *CoinSwitch) Rate(depositCoin, destinationCoin string) (*Rate, error) {
	path := "/rate"
	result := new(RateResult)

	_, err := c.client.R().
		SetResult(result).
		SetBody(map[string]string{
			"depositCoin":     depositCoin,
			"destinationCoin": destinationCoin,
		}).
		Post(path)

	if err != nil {
		return nil, err
	}

	if result.Data == nil {
		return nil, errors.New("trade_pair_unavailable")
	}

	return result.Data, nil
}

// Order - Create and Order
// https://developer.coinswitch.co/reference#apiorder
func (c *CoinSwitch) CreateOrder(request *OrderRequest) (*Order, error) {
	path := "/order"
	result := new(OrderResult)

	_, err := c.client.R().
		SetResult(result).
		SetBody(request).
		Post(path)

	if err != nil {
		return nil, err
	}

	return &result.Data, nil
}

// GetOrder - Get status for an Order ID
// https://developer.coinswitch.co/reference#apiorderstatus
func (c *CoinSwitch) GetOrderStatus(orderID string) (*OrderStatus, error) {
	path := "/order/" + orderID
	result := new(OrderStatusResult)

	_, err := c.client.R().
		SetResult(result).
		Get(path)

	if err != nil {
		return nil, err
	}

	return &result.Data, nil
}

// GetOrders - Get all orders created with your API Key
// https://developer.coinswitch.co/reference#get-v2orders
func (c *CoinSwitch) GetOrders(start, count int) ([]Order, error) {
	path := "/orders"
	result := new(OrdersResult)

	_, err := c.client.R().
		SetResult(result).
		SetQueryParam("start", strconv.Itoa(start)).
		SetQueryParam("count", strconv.Itoa(count)).
		Get(path)

	if err != nil {
		return nil, err
	}

	return result.Data.Items, nil
}
