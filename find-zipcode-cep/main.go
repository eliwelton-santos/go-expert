package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
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
	for _, zipcode := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + zipcode + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Response read error: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse response error: %v\n", err)
		}

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Create file error: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("Zipcode: %s, City: %s, UF: %s", data.Cep, data.Localidade, data.Uf))

		fmt.Println("File created successfully!")
		fmt.Println("City:", data.Localidade)
	}
}
