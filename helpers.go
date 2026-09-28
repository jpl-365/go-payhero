package gopayhero

import (
	"encoding/base64"
	"fmt"
)

func GetToken(apiUsername string, apiPassword string) string {
	credentials := fmt.Sprintf("%s:%s", apiUsername, apiPassword)
	return base64.StdEncoding.EncodeToString([]byte(credentials))

}
