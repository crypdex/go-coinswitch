package coinswitch

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
)

var baseURL = "https://api.coinswitch.co/v2"

// ChangeNow is the root client object
type CoinSwitch struct {
	client *resty.Client
}

// New creates a configured instance of ChangeNow
func New(apikey string) *CoinSwitch {
	client := resty.New().
		// SetHTTPMode().
		SetHostURL(baseURL).
		SetHeader("x-api-key", apikey).
		SetHeader("x-user-ip", "172.98.82.16").
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(1)).
		OnAfterResponse(checkResponse)

	return &CoinSwitch{client}
}

// checkResponse for errors with the request and return them
func checkResponse(client *resty.Client, response *resty.Response) error {
	result := new(Result)
	fmt.Println(string(response.Body()))
	if err := json.Unmarshal(response.Body(), result); err != nil {
		return err
	}
	if result.Success {
		return nil
	}

	// There was an error
	return fmt.Errorf("coinswitch error: %s %s", result.Code, result.Error)
}
