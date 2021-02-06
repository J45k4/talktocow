package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/auth"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

type ErrorCode uint32

const (
	InvalidCredentials  ErrorCode = 9000
	InternalServerError           = 9001
)

type Error struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func CreateErrorResponse(message string, code ErrorCode) ErrorResponse {
	return ErrorResponse{
		Error: Error{
			Message: message,
			Code:    code,
		},
	}
}

func HandleLogin(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	var loginRequest LoginRequest

	ctx.BindJSON(&loginRequest)

	fmt.Printf("New login attempt from {%s}\n", loginRequest.Username)

	var user models.User

	err := models.Users(
		qm.Where("username = ?", loginRequest.Username),
	).Bind(ctx.Request.Context(), db, &user)

	if err != nil {
		fmt.Println("User not found")

		ctx.Status(http.StatusForbidden)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse("Credentials are incorrect", InvalidCredentials))

		return
	}

	if auth.CheckPasswordHash(loginRequest.Password, user.PasswordHash.String) == false {
		fmt.Println("Password is incorrect")

		ctx.Status(http.StatusForbidden)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse("Credentials are incorrect", InvalidCredentials))

		return
	}

	token, err := auth.GenerateTokenFromObject(UserSession{
		UserID:   int32(user.ID),
		UserName: user.Name.String,
	})

	if err != nil {
		fmt.Println("Generating token failed")

		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse("Internal server error", InternalServerError))
		ctx.Abort()

		return
	}

	resp := LoginResponse{
		Token:    string(token),
		UserID:   fmt.Sprint(user.ID),
		Username: user.Name.String,
	}

	ctx.JSON(200, resp)
}
