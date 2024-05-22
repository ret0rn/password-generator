package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type requestBody struct {
	Length  uint   `json:"length" validate:"min=0,max=256"`
	Symbols string `json:"symbols" validate:"min=0,max=256"`
}

type responseBody struct {
	Result string `json:"result"`
}

const dict = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func generator(length uint, symbols string) string {
	password := make([]byte, length)
	for i := range password {
		password[i] = dict[rand.Intn(len(dict))] // #nosec G404
	}
	if symbols != "" {
		for i := 0; i < rand.Intn(len(password)); i++ { // #nosec G404
			password[rand.Intn(len(password))] = symbols[rand.Intn(len(symbols))] // #nosec G404
		}
	}
	return string(password)
}

func GeneratePassword(w http.ResponseWriter, r *http.Request) {
	var (
		req requestBody
		err error
	)
	if r != nil && r.Body != http.NoBody {
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			logrus.Errorf("Decode request error %v", err)
			http.Error(w, "cant's get request data", http.StatusBadRequest)
			return
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			logrus.Errorf("Validation request error %v", err)
			http.Error(w, fmt.Sprintf("cant's validate request data %s", err.Error()), http.StatusBadRequest)
			return
		}
	}

	if req.Length == 0 {
		req.Length = 12
	}

	var res responseBody
	res.Result = generator(req.Length, req.Symbols)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		logrus.Errorf("Encode response error %v", err)
		return
	}
}
