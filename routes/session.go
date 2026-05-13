package routes

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/auth"
	"github.com/j45k4/talktocow/config"
)

const AuthCookieName = "talktocow_auth"

const authCookieMaxAge = int((30 * 24 * time.Hour) / time.Second)

func authCookieSecure(ctx *gin.Context) bool {
	if config.AuthCookieSecure {
		return true
	}

	if ctx.Request.TLS != nil {
		return true
	}

	forwardedProto := strings.ToLower(ctx.GetHeader("x-forwarded-proto"))

	if strings.Contains(forwardedProto, "https") {
		return true
	}

	return strings.EqualFold(ctx.GetHeader("x-forwarded-ssl"), "on")
}

func SetAuthCookie(ctx *gin.Context, token string) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     AuthCookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   authCookieMaxAge,
		HttpOnly: true,
		Secure:   authCookieSecure(ctx),
		SameSite: http.SameSiteLaxMode,
	})
}

func ClearAuthCookie(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     AuthCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   authCookieSecure(ctx),
		SameSite: http.SameSiteLaxMode,
	})
}

func bearerTokenFromHeader(authHeader string) string {
	parts := strings.Fields(authHeader)

	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}

	return parts[1]
}

func authenticateToken(token string, out *UserSession) error {
	return auth.DecodeObjectFromToken(token, out)
}

type authTokenSource struct {
	name  string
	token string
}

func getUserSessionAndTokenSourceFromRequest(ctx *gin.Context) (UserSession, authTokenSource, bool) {
	sources := []authTokenSource{
		{name: "authorization header", token: bearerTokenFromHeader(ctx.GetHeader("authorization"))},
	}

	if cookie, err := ctx.Request.Cookie(AuthCookieName); err == nil {
		sources = append(sources, authTokenSource{name: "auth cookie", token: cookie.Value})
	}

	sources = append(sources, authTokenSource{name: "query token", token: ctx.Query("token")})

	for _, source := range sources {
		if source.token == "" {
			continue
		}

		var userSession UserSession

		if err := authenticateToken(source.token, &userSession); err != nil {
			log.Printf("auth token from %s is invalid: %v", source.name, err)
			continue
		}

		return userSession, source, true
	}

	return UserSession{}, authTokenSource{}, false
}

func GetUserSessionFromRequest(ctx *gin.Context) (UserSession, bool) {
	userSession, _, ok := getUserSessionAndTokenSourceFromRequest(ctx)
	return userSession, ok
}

func SessionMiddleware(ctx *gin.Context) {
	userSession, tokenSource, ok := getUserSessionAndTokenSourceFromRequest(ctx)

	if !ok {
		log.Println("User is not authorized")

		ctx.Status(http.StatusUnauthorized)

		ctx.Abort()
		return
	}

	if tokenSource.name == "authorization header" {
		SetAuthCookie(ctx, tokenSource.token)
	}

	ctx.Set("userSession", userSession)
}
