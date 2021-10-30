package routes

import (
	"fmt"
	"log"
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

func HandleLogin(ctx *gin.Context) {
	db := GetDBFromContext(ctx)

	var loginRequest LoginRequest

	ctx.BindJSON(&loginRequest)

	fmt.Printf("New login attempt from {%s}\n", loginRequest.Username)

	user, err := models.Users(
		qm.Where("username = ?", loginRequest.Username),
	).One(ctx.Request.Context(), db)

	if err != nil {
		log.Printf("fetching user failed %v", err)

		ctx.JSON(http.StatusInternalServerError, CreateErrorResponse(InternalServerError, ""))
		return
	}

	if user == nil {
		fmt.Println("User not found")

		ctx.Status(http.StatusForbidden)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidCredentials, "Credentials are incorrect"))
		return
	}

	if !auth.CheckPasswordHash(loginRequest.Password, user.PasswordHash.String) {
		fmt.Println("Password is incorrect")

		ctx.Status(http.StatusForbidden)
		ctx.JSON(http.StatusForbidden, CreateErrorResponse(InvalidCredentials, "Credentials are incorrect"))

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
