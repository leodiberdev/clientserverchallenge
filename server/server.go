package server

import (
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

func RequestRates() (Response, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	req, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return Response{}, err
	}

	var rateData Response
	err = json.Unmarshal(res, &rateData)
	if err != nil {
		return Response{}, err
	}

	return rateData, nil
}
