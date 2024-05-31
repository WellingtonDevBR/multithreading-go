package api

import (
	"cep-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchFromViaCEP(cep string, ch chan<- models.CepResponse) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		ch <- models.CepResponse{Source: "ViaCEP", Localidade: "Error"}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- models.CepResponse{Source: "ViaCEP", Localidade: "Error"}
		return
	}
	var data models.CepResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		ch <- models.CepResponse{Source: "ViaCEP", Localidade: "Error"}
		return
	}
	data.Source = "ViaCEP"
	ch <- data
}
