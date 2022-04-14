package routes

import (
	"database/sql"
	"strconv"

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

func GetOffsetAndLimit(ctx *gin.Context, defOffset int, defLim int) (int, int) {
	offset := defOffset
	limit := defLim

	if ctx.Query("offset") != "" {
		offset, _ = strconv.Atoi(ctx.Query("offset"))
	}

	if ctx.Query("limit") != "" {
		limit, _ = strconv.Atoi(ctx.Query("limit"))
	}

	return offset, limit
}

func getNumParam(ctx *gin.Context, paramName string) int {
	p := ctx.Param(paramName)

	numP := 0

	if p != "" {
		numP, _ = strconv.Atoi(p)
	}

	return numP
}
