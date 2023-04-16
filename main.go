package main

import (
	"github.com/leonardodiber/clientserverchallenge/client"
	"github.com/leonardodiber/clientserverchallenge/server"
)

func main() {
	server.StartServer()
	client.RequestBid()

	select {}
}
