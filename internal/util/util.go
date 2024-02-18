package util

import (
	"encoding/json"
	"net/http"
)

// Reads request json body into the data parameter
func ReadRequestBody(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// Prevent large request from reaching server
	maxBytes := 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	// Ensure requests have valid bodies. If not return 400 Bad Request
	dec.DisallowUnknownFields()
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// Sets response status code and writes json data from data paramater onto response
func WriteJSON(w http.ResponseWriter, data interface{}, status int) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		return err
	}

	return nil
}
