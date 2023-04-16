package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	UsdBrlRate UsdBrlRate `json:"USDBRL"`
}

type UsdBrlRate struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func RequestRates(ctx context.Context) (Response, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return Response{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	var rateData Response
	err = json.Unmarshal(responseData, &rateData)
	if err != nil {
		return Response{}, err
	}

	return rateData, nil
}
