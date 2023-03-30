package routes

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/eventbus"
)

func GetDBFromContext(ctx *gin.Context) *sql.DB {
	db, _ := ctx.Get("db")

	db2 := db.(*sql.DB)

	return db2
}

func GetEventbusFromContext(ctx *gin.Context) *eventbus.Eventbus {
	e, _ := ctx.Get("eventbus")

	e2 := e.(*eventbus.Eventbus)

	return e2
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

func getNumParam(ctx *gin.Context, paramName string) (int, error) {
	p := ctx.Param(paramName)

	numP, err := strconv.Atoi(p)

	return numP, err
}
