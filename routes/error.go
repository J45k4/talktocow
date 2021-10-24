package routes

type ErrorCode string

const (
	InvalidCredentials  ErrorCode = "INVALID_CREDENTIALS"
	InternalServerError           = "INTERNAL_SERVER_ERROR"
	InvalidInput                  = "INVALID_INPUT"
)

type Error struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func CreateErrorResponse(code ErrorCode, message string) ErrorResponse {
	return ErrorResponse{
		Error: Error{
			Message: message,
			Code:    code,
		},
	}
}
