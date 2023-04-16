package main

import (
	"context"
	"fmt"
	"time"

	"github.com/leonardodiber/clientserverchallenge/server"
)

func main() {
	reqCtx, cancelReq := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelReq()
	rates, _ := server.RequestRates(reqCtx)

	fmt.Println(rates.UsdBrlRate.Bid)
}
