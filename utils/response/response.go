package response

import (
	"encoding/json"
	"net/http"
	"reflect"
)

const (
	// RecordNotFound : Code error for not found
	RecordNotFound = 20023
	// ResourceNotFound : Code error for recurse not found
	ResourceNotFound = 20022
	// BadRequest : Code error for invalid request
	BadRequest = 20001
	// InternalServerError : Code error for internal server error
	InternalServerError = 10000
	// Unauthorized : Code error for request unauthorized
	Unauthorized = 30001
)

// SuccessResponse : Struct of success for response
type SuccessResponse struct {
	Meta    meta        `json:"meta"`
	Records interface{} `json:"records"`
}

// ErrorResponse : Struct of error for response
type ErrorResponse struct {
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
	ErrorCode        int    `json:"errorCode"`
	MoreInfo         string `json:"moreInfo"`
}

type meta struct {
	Server      string `json:"server"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
	RecordCount int    `json:"recordCount"`
}

// GenerateSuccessResponse : Generate response of success
func GenerateSuccessResponse(obj interface{}, limit, offset, recordCount int) SuccessResponse {
	var successResponse SuccessResponse

	successResponse.Meta.Server = "localhost"
	successResponse.Meta.Limit = limit
	successResponse.Meta.Offset = offset
	successResponse.Meta.RecordCount = recordCount

	if reflect.TypeOf(obj).Kind() != reflect.Slice {
		records := make([]interface{}, 1)
		records[0] = obj
		successResponse.Records = records
		return successResponse
	}

	successResponse.Records = obj

	return successResponse
}

// GenerateErrorResponse : Generate error response
func GenerateErrorResponse(errorCode int, developerMessage, userMessage string) ErrorResponse {
	return ErrorResponse{
		DeveloperMessage: developerMessage,
		ErrorCode:        errorCode,
		UserMessage:      userMessage,
		MoreInfo:         "http://www.developer.apiluiza.com.br/errors",
	}
}

// GenerateHTTPResponse : Generate http response
func GenerateHTTPResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
