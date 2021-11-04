package routes

type ErrorCode string

const (
	InvalidCredentials  ErrorCode = "INVALID_CREDENTIALS"
	InternalServerError ErrorCode = "INTERNAL_SERVER_ERROR"
	InvalidInput        ErrorCode = "INVALID_INPUT"
)

type Error struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func CreateErrorResponse(code ErrorCode, message string) ErrorResponse {
	msg := message

	if msg == "" {
		msg = string(code)
	}

	return ErrorResponse{
		Error: Error{
			Message: msg,
			Code:    code,
		},
	}
}
