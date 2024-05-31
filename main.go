package main

import (
	"cep-api/api"
	"cep-api/models"
	"fmt"
	"time"
)

func main() {
	cep := "01153000"
	ch := make(chan models.CepResponse)

	go api.FetchFromBrasilAPI(cep, ch)
	go api.FetchFromViaCEP(cep, ch)

	select {
	case res := <-ch:
		if res.Localidade == "Error" {
			fmt.Println("Erro ao buscar o CEP.")
		} else {
			fmt.Printf("Resposta da API %s:\nCep: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
				res.Source, res.Cep, res.Logradouro, res.Complemento, res.Bairro, res.Localidade, res.Uf, res.Ibge, res.Gia, res.Ddd, res.Siafi)
		}
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: timeout ao buscar o CEP.")
	}
}
