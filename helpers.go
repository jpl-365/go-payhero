package gopayhero

import (
	"encoding/base64"
	"fmt"
)

func getToken(apiUsername string, apiPassword string) string {
	credentials := fmt.Sprintf("%s:%s", apiUsername, apiPassword)
	return base64.StdEncoding.EncodeToString([]byte(credentials))

}

func (c *Client) basicAuth() string {
	return "Basic " + getToken(c.apiUsername, c.apiPassword)
}
