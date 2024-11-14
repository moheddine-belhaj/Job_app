package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody parses the request body into the given struct and returns an error if parsing fails.
func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, x); err != nil {
		return err
	}
	return nil
}
