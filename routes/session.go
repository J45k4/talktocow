package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/auth"
)

func SessionMiddleware(ctx *gin.Context) {
	var token string

	authHeader := ctx.GetHeader("authorization")

	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		query := ctx.Request.URL.Query()

		token = query.Get("token")
	}

	if token == "" {
		log.Println("User is not authorized")

		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
		return
	}

	var userSession UserSession

	err := auth.DecodeObjectFromToken(token, &userSession)

	if err != nil {
		db := GetDBFromContext(ctx)
		apiKeySession, ok := AuthenticateAPIKey(ctx, db, token)
		if !ok {
			log.Println("User token is invalid", err)

			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		userSession = apiKeySession
	}

	ctx.Set("userSession", userSession)

	log.Println("userSession", userSession)
}
