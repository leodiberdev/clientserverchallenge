package main

import (
	"fmt"

	"github.com/leonardodiber/clientserverchallenge/server"
)

func main() {
	rates, _ := server.RequestRates()

	fmt.Println(rates.UsdBrlRate.Bid)
}
