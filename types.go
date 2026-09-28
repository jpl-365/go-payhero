package gopayhero

import "net/http"

// configuration

type Client struct {
	apiUsername string
	apiPassword string
	baseURL     string
	httpClient  *http.Client
}
