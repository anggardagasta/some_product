package response

import (
	"encoding/json"
	"net/http"
)

const (
	MessageSucceed       = "Succeed"
	MessageUserNotFount  = "User Not Found"
	MessageInternalError = "Internal Error"
	MessageBadRequest    = "Bad Request"
)

type (
	responseBody struct {
		Status  int64       `json:"status"`
		Message string      `json:"message"`
		Error   string      `json:"error"`
		Data    interface{} `json:"data,omitempty"`
	}
)

func Result(w http.ResponseWriter, message string, statusCode int64) {
	result := responseBody{}
	result.Status = statusCode
	result.Message = message

	_ = json.NewEncoder(w).Encode(result)
}

func ResultWithData(w http.ResponseWriter, data interface{}, message string, statusCode int64) {
	result := responseBody{}
	result.Status = statusCode
	result.Message = message
	result.Data = data

	_ = json.NewEncoder(w).Encode(result)
}

func ResultError(w http.ResponseWriter, statusCode int64, message string, err error) {
	result := responseBody{}
	result.Status = statusCode
	result.Message = message
	result.Error = err.Error()
	_ = json.NewEncoder(w).Encode(result)
}
