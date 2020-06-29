package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type (
	ExchangeRate struct {
		Last struct {
			Ask float64 `json:"ask"`
		} `json:"last"`
		Message string `json:"message,omitempty"`
	}
)

func TestParse(t *testing.T) {
	url := "https://api.polygon.io/v1/last_quote/currencies/JPY/USD?amount=1&precision=2&apiKey=redacted"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[ERROR] unable to get quote from polygon.io API , error : %v", err)

	}

	jsonResp, _ := ioutil.ReadAll(resp.Body)

	//s := `{"last":{"ask":1.25638,"bid":1.25612,"exchange":48,"timestamp":1592356494000},"status":"success","symbol":"GBP/USD"}`

	ex := ExchangeRate{}
	err = json.Unmarshal(jsonResp, &ex)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(ex)
	}

}
