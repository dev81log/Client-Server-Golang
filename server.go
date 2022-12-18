package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CotacaoMoeda struct {
	Usdbrl struct {
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
	} `json:"USDBRL"`
}

func main() {
	db, err := sql.Open("sqlite3", "../db/cotacao.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	cotacao, err := apiCotacao(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = insertCotacao(ctx, db, cotacao)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Dólar : %s", cotacao.Usdbrl.Bid)
		_, err := apiCotacao(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}

func apiCotacao(ctx context.Context) (*CotacaoMoeda, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	client := http.Client{}
	request, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var cotacao *CotacaoMoeda
	if err := json.Unmarshal(body, &cotacao); err != nil {
		return nil, err
	}
	return cotacao, nil
}

func insertCotacao(ctx context.Context, db *sql.DB, cotacao *CotacaoMoeda) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare("insert into dolar (cotacao) values (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		cotacao.Usdbrl.Bid,
	)
	if err != nil {
		return err
	}
	return nil
}
