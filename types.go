package gopayhero

import (
	"net/http"
	"time"
)

// configuration

type Client struct {
	apiUsername string
	apiPassword string
	baseURL     string
	httpClient  *http.Client
}

// payment channels
type PaymentChannelsResponse struct {
	PaymentChannels []PaymentChannel `json:"payment_channels"`
	Pagination      Pagination       `json:"pagination"`
}

type PaymentChannel struct {
	ID              int64     `json:"id"`
	TransactionType string    `json:"transaction_type"`
	ChannelType     string    `json:"channel_type"`
	AccountID       int64     `json:"account_id"`
	ShortCode       string    `json:"short_code"`
	AccountNumber   string    `json:"account_number"`
	Description     string    `json:"description"`
	IsActive        bool      `json:"is_active"`
	BalancePlain    *float64  `json:"balance_plain"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Pagination struct {
	Count    int  `json:"count"`
	NextPage *int `json:"next_page"`
	NumPages int  `json:"num_pages"`
	Page     int  `json:"page"`
	Per      int  `json:"per"`
	PrevPage *int `json:"prev_page"`
}

type ErrorMessage struct {
	ErrorMessage string `json:"error_message"`
}
