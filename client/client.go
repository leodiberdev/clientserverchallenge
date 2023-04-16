package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func RequestBid() error {
	url := "http://localhost:8080/cotacao"

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var bid string
	err = json.Unmarshal(responseData, &bid)
	if err != nil {
		return err
	}

	cotacao := fmt.Sprintf("Dolar: %s", bid)

	err = ioutil.WriteFile("cotacao.txt", []byte(cotacao), 0644)

	fmt.Println(bid)
	return nil
}
