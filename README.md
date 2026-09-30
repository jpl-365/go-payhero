# Golang Library for Payhero API.

## Usage

```
package main 

import "github.com/jpl-365/go-payhero"

func main(){
    client:=go-payhero.NewClient(
        os.GetEnv(YOUR_API_USERNAME),
        os.GetEnv(YOUR_API_PASSWORD),
    )

    // Get Payment Channels
    channels, err := gopayhero.GetPaymentChannels(ctx context.Context())
    if err !=nil{
        panic(err)
    }

}

```

## Features 

### payment channels

- Get payment channels 



