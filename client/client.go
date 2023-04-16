package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	file, err := os.OpenFile("cotacao.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Um erro ocorreu ao tentar abrir/criar o arquivo")
		return err
	}
	defer file.Close()

	_, err = file.WriteString(cotacao)
	if err != nil {
		fmt.Println("Um erro ocorreu ao tentar salvar os dados no arquivo")
		return err
	}

	fmt.Printf("Cotacao salva no arquivo com sucesso!\n")

	return nil
}
