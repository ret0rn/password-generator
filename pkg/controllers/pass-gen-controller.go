package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type requestBody struct {
	Length  uint   `json:"length"`
	Symbols string `json:"symbols"`
}

type responseBody struct {
	Result string `json:"result"`
}

const dict = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GeneratePassword(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Length == 0 {
		req.Length = 12
	}
	password := make([]byte, req.Length)
	for i := range password {
		password[i] = dict[rand.Intn(len(dict))]
	}
	symbols := req.Symbols
	if symbols != "" {
		for i := 0; i < rand.Intn(len(password)); i++ {
			password[rand.Intn(len(password))] = symbols[rand.Intn(len(symbols))]
		}
	}

	var res responseBody
	res.Result = string(password)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
