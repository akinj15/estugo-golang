package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	urlBase := "https://viacep.com.br/ws/"
	urlFinal := "/json/"
	for _, cep := range os.Args[1:] {
		url := urlBase + cep + urlFinal
		req, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		var cepConsultado ViaCep

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(res, &cepConsultado)
		if err != nil {
			panic(err)
		}
		println(url)
		println(cepConsultado.Logradouro)

		file, err := os.Create("consultaCep.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.Write(res)
	}
}
