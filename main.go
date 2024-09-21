package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Address struct {
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

type Result struct {
	Source  string
	Address Address
	Error   error
}

func fetchFromBrasilAPI(cep string, ch chan<- Result) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{Source: "BrasilAPI", Error: err}
		return
	}
	defer resp.Body.Close()

	var address Address
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		ch <- Result{Source: "BrasilAPI", Error: err}
		return
	}

	ch <- Result{Source: "BrasilAPI", Address: address}
}

func fetchFromViaCEP(cep string, ch chan<- Result) {

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{Source: "ViaCEP", Error: err}
		return
	}
	defer resp.Body.Close()

	var address Address
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		ch <- Result{Source: "ViaCEP", Error: err}
		return
	}

	ch <- Result{Source: "ViaCEP", Address: address}
}

func main() {
	cep := "01153000"
	ch := make(chan Result, 2)

	go fetchFromBrasilAPI(cep, ch)
	go fetchFromViaCEP(cep, ch)

	select {
	case result := <-ch:
		if result.Error != nil {
			fmt.Printf("Error fetching from %s: %v\n", result.Source, result.Error)
		} else {
			fmt.Printf("Result from %s: %+v\n", result.Source, result.Address)
		}
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No response within 1 second")
	}
} //TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
