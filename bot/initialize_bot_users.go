package bot

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InitializeBot(name string, db *sql.DB, ctx context.Context) {
	user, err := models.Users(qm.Where("name = ?", name)).One(ctx, db)

	if err == nil || user != nil {
		return
	}

	user = &models.User{
		Name: null.NewString(name, true),
		Bot:  true,
	}

	err = user.Insert(ctx, db, boil.Infer())

	if err != nil {
		panic(err)
	}
}

func InitializeBots(db *sql.DB) {
	InitializeBot("CowGPT", db, context.Background())
}

func GetCowGPTUser(db *sql.DB) *models.User {
	ctx := context.Background()

	user, err := models.Users(qm.Where("name = ?", "CowGPT")).One(ctx, db)

	if err != nil {
		fmt.Printf("error finding user: %v", err)

		return nil
	}

	return user
}
