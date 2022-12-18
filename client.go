package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	buscarCotacao(ctx)
}

func buscarCotacao(ctx context.Context) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	client := http.Client{}
	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := file.Write(body)
	if err != nil {
		panic(err)
	}

	println(n, "Cotação salva com sucesso em cotacao.txt")
	file.Close()

	return body, nil
}
