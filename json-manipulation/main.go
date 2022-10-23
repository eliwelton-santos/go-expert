package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	// Encoder
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	// Unmarshal
	pureJson := []byte(`{"number":2,"balance":200}`)
	var accountX Account
	err = json.Unmarshal(pureJson, &accountX)
	if err != nil {
		println(err)
	}
	println(accountX.Balance)
}
