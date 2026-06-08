package shared

import (
	"encoding/json"
	"net/http"
)

// ErrorCode represents a stable error code
type ErrorCode string

const (
	CATALOG_UNAVAILABLE   ErrorCode = "CATALOG_UNAVAILABLE"
	INVALID_QUANTITY      ErrorCode = "INVALID_QUANTITY"
	PRODUCT_NOT_FOUND     ErrorCode = "PRODUCT_NOT_FOUND"
	PRODUCT_UNAVAILABLE   ErrorCode = "PRODUCT_UNAVAILABLE"
	PRODUCT_OUT_OF_STOCK  ErrorCode = "PRODUCT_OUT_OF_STOCK"
	INSUFFICIENT_STOCK    ErrorCode = "INSUFFICIENT_STOCK"
	CART_UPDATE_FAILED    ErrorCode = "CART_UPDATE_FAILED"
)

// APIError represents an error response
type APIError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// ErrorResponse wraps APIError for response
type ErrorResponse struct {
	Error APIError `json:"error"`
}

// SuccessResponse wraps data for response
type SuccessResponse struct {
	Data interface{} `json:"data,omitempty"`
}

// WriteErrorResponse writes error to response writer
func WriteErrorResponse(w http.ResponseWriter, statusCode int, code ErrorCode, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: APIError{
			Code:    code,
			Message: message,
		},
	})
}

// WriteSuccessResponse writes success data to response writer
func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
