package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetDBFromContext(ctx *gin.Context) *sql.DB {
	db, _ := ctx.Get("db")

	db2 := db.(*sql.DB)

	return db2
}

func GetUserSessionFromContext(ctx *gin.Context) UserSession {
	userSession, _ := ctx.Get("userSession")

	u := userSession.(UserSession)

	return u
}

func DoesMaskHaveField(mask []string, fieldName string) bool {
	for _, m := range mask {
		if m == fieldName {
			return true
		}
	}

	return false
}
