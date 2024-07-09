package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NeiderFajardo/pkg/validator"
)

func Encode[T any](w http.ResponseWriter, status int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func Decode[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode json: %w", err)
	}
	return data, nil
}

func DecodeValid[T validator.Validator](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode json: %w", err)
	}
	if problems := data.Valid(r.Context()); len(problems) > 0 {
		for _, problem := range problems {
			return data, &problem
		}
	}
	return data, nil
}
