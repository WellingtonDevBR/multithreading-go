package api

import (
	"cep-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BrasilAPIResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"street"`
	Complemento string `json:"complement"`
	Bairro      string `json:"neighborhood"`
	Localidade  string `json:"city"`
	Uf          string `json:"state"`
}

func FetchFromBrasilAPI(cep string, ch chan<- models.CepResponse) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		ch <- models.CepResponse{Source: "BrasilAPI", Localidade: "Error"}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- models.CepResponse{Source: "BrasilAPI", Localidade: "Error"}
		return
	}
	var data BrasilAPIResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		ch <- models.CepResponse{Source: "BrasilAPI", Localidade: "Error"}
		return
	}
	ch <- models.CepResponse{
		Cep:         data.Cep,
		Logradouro:  data.Logradouro,
		Complemento: data.Complemento,
		Bairro:      data.Bairro,
		Localidade:  data.Localidade,
		Uf:          data.Uf,
		Source:      "BrasilAPI",
	}
}
