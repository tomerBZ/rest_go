package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

func ToJson(data interface{}, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(data)
	return
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
