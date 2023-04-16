package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"time"

	// "log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
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

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	rates, err := RequestRates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = saveRates(rates.UsdBrlRate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response, err := json.Marshal(rates.UsdBrlRate.Bid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func StartServer() {
	http.HandleFunc("/cotacao", handleCotacao)
	http.ListenAndServe(":8080", nil)
}

func RequestRates() (Response, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

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

func saveRates(rates UsdBrlRate) error {
	// sets up the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// open the db connection
	db, err := sql.Open("sqlite3", "./db/mydb.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// creates the table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS UsdBrlRate (
        Code TEXT,
        Codein TEXT,
        Name TEXT,
        High TEXT,
        Low TEXT,
        VarBid TEXT,
        PctChange TEXT,
        Bid TEXT,
        Ask TEXT,
        Timestamp TEXT,
        CreateDate TEXT
    );`

	err = createTable(ctx, db, createTableSQL)
	if err != nil {
		return err
	}

	// prepare the statement to insert the data into the table
	insertSQL := `INSERT INTO UsdBrlRate (Code, Codein, Name, High, Low, VarBid, PctChange, Bid, Ask, Timestamp, CreateDate) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}

	// insert the data into the table
	_, err = stmt.ExecContext(
		ctx,
		rates.Code,
		rates.Codein,
		rates.Name,
		rates.High,
		rates.Low,
		rates.VarBid,
		rates.PctChange,
		rates.Bid,
		rates.Ask,
		rates.Timestamp,
		rates.CreateDate,
	)
	if err != nil {
		return err
	}

	return nil
}

func createTable(ctx context.Context, db *sql.DB, query string) error {
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
