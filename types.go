package coinswitch

// Error Codes
// -----------
// authentication_error
// invalid_address
// invalid_input
// invalid_symbol
// temp_inactive_coin
// inactive_coin
// trade_pair_disabled
// limit_exceeded
// unknown
// min_limit_breached
// max_limit_breached

type Result struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Error   string `json:"error"`
}

type Coin struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	IsActive bool   `json:"isActive"`
	IsFiat   bool   `json:"isFiat"`
	LogoURL  string `json:"logoUrl"`
}

type CoinResult struct {
	Result
	Data []Coin `json:"data"`
}

type Pair struct {
	DepositCoin     string `json:"depositCoin"`
	DestinationCoin string `json:"destinationCoin"`
	IsActive        bool   `json:"isActive"`
}

type PairsResult struct {
	Result
	Data []Pair `json:"data"`
}

type Rate struct {
	Rate                    float64 `json:"rate"`
	MinerFee                float64 `json:"miner_fee"`
	LimitMinDepositCoin     float64 `json:"limitMinDepositCoin"`
	LimitMaxDepositCoin     float64 `json:"limitMaxDepositCoin"`
	LimitMinDestinationCoin float64 `json:"limitMinDestinationCoin"`
	LimitMaxDestinationCoin float64 `json:"limitMaxDestinationCoin"`
}

type RateResult struct {
	Result
	Data *Rate `json:"data"`
}

type Order struct {
	OrderID                       string  `json:"orderId"`
	ExchangeAddress               Address `json:"exchangeAddress"`
	ExpectedDepositCoinAmount     float64 `json:"expectedDepositCoinAmount"`
	ExpectedDestinationCoinAmount float64 `json:"expectedDestinationCoinAmount"`
}

type OrderResult struct {
	Result
	Data Order `json:"data"`
}

type Address struct {
	Address string  `json:"address"`
	Tag     *string `json:"tag"`
}

type OrderRequest struct {
	DepositCoin           string   `json:"depositCoin"`                     // Coin that will be deposited to CoinSwitch
	DestinationCoin       string   `json:"destinationCoin"`                 // Coin that CoinSwitch will convert the depositCoin to
	DepositCoinAmount     float64  `json:"depositCoinAmount"`               // Amount of depositCoin that will be sent to CoinSwitch
	DestinationCoinAmount *float64 `json:"destinationCoinAmount,omitempty"` // Amount of destinationCoin required from CoinSwitch
	DestinationAddress    Address  `json:"destinationAddress"`              // JSON FIELD - Address to which destinationCoin will be sent
	RefundAddress         *Address `json:"refundAddress,omitempty"`         // JSON FIELD - Address to which depositCoin will be refunded in case of unforeseen errors
}

type OrderStatus struct {
	OrderID               string  `json:"orderId"`
	exchangeAddress       Address `json:"exchangeAddress"`
	destinationAddress    Address `json:"destinationAddress"`
	Status                string  `json:"status"`
	CreatedAt             uint    `json:"createdAt"`
	InputTransactionHash  string  `json:"inputTransactionHash"`
	OutputTransactionHash string  `json:"outputTransactionHash"`
	DepositCoin           string  `json:"depositCoin"`
	DestinationCoin       string  `json:"destinationCoin"`
	DepositCoinAmount     float64 `json:"depositCoinAmount"`
	DestinationCoinAmount float64 `json:"destinationCoinAmount"`
	ValidTill             *string `json:"validTill"`
	UserReferenceId       string  `json:"userReferenceId"`
}

type OrderStatusResult struct {
	Result
	Data OrderStatus `json:"data"`
}

type OrdersResult struct {
	Result
	Data struct {
		Count int     `json:"count"`
		Items []Order `json:"items"`
	} `json:"data"`
}
