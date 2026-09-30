package gopayhero

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var channelUrl = "/payment_channels?is_active=true"

func (c *Client) GetPaymentChannels(ctx context.Context) (PaymentChannelsResponse, error) {
	var channels PaymentChannelsResponse
	var error ErrorMessage
	client := c.httpClient
	req, _ := http.NewRequest("GET", c.baseURL+channelUrl, nil)
	req.Header.Set("Authorization", c.basicAuth())

	resp, err := client.Do(req)
	if err != nil {
		return PaymentChannelsResponse{}, errors.New("failed to send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		json.NewDecoder(resp.Body).Decode(&error)
		return PaymentChannelsResponse{}, errors.New(error.ErrorMessage)

	}

	json.NewDecoder(resp.Body).Decode(&channels)
	return channels, nil

}
