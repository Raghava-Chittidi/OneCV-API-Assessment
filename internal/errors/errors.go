package errors

import (
	"Raghava/OneCV-Assignment/internal/api"
	"Raghava/OneCV-Assignment/internal/util"
	"fmt"
	"net/http"
)

var (
	NotFoundErrorString = "not found"
	UnprocessableEntityErrorString = "unprocessable entity"
	BadRequestErrorString = "bad request"
	InternalServerErrorString = "internal server error"

)

// Get error message based on the status code given
func GetErrorMessage(err error, status int) error {
	switch status {
	case http.StatusNotFound:
		return fmt.Errorf("error %s: %w", NotFoundErrorString, err)
	case http.StatusUnprocessableEntity:
		return fmt.Errorf("error %s: %w", UnprocessableEntityErrorString, err)
	case http.StatusBadRequest:
		return fmt.Errorf("error %s: %w", BadRequestErrorString, err)
	default:
		return fmt.Errorf("error %s: %w", InternalServerErrorString, err)
	}
}

// Format error with gorm model name
func FormatError(modelName string, err error) error {
	return fmt.Errorf("%s %w", modelName, err)
}

// Sets json error message and response status code
func ErrorJSON(w http.ResponseWriter, err error, status int) {
	errMessage := GetErrorMessage(err, status).Error()
	res := api.Response{Message: errMessage}
	util.WriteJSON(w, res, status)
}