package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/j45k4/talktocow/config"
	"github.com/j45k4/talktocow/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	connectionString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", config.DBName, config.DBUser, config.DBPassword, config.DbHost, config.DBPort)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Sql open error ", err)

		return
	}

	ctx := context.Background()

	messages, _ := models.Messages().All(ctx, db)

	for _, message := range messages {
		if message.Reference.Valid == false {
			message.Reference = null.StringFrom(uuid.New().String())

			message.Update(ctx, db, boil.Infer())
		}
	}
}
