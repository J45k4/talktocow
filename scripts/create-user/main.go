package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/j45k4/talktocow/auth"
	"github.com/j45k4/talktocow/config"
	"github.com/j45k4/talktocow/models"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 2 {
		panic("Incorrect number of arguments")
	}

	username := argsWithoutProg[0]
	password := argsWithoutProg[1]

	fmt.Printf("Creating user username %v password %v\n", username, password)

	connectionString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", config.DBName, config.DBUser, config.DBPassword, config.DbHost, config.DBPort)

	fmt.Println("Connection string:", connectionString)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println("Sql open error ", err)

		return
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)

	if err != nil {
		fmt.Printf("err %v \n", err)

		panic("Failed to execute migrations")
	}

	ctx := context.Background()

	pass, err := auth.HashPassword(password)

	if err != nil {
		fmt.Printf("Password hashing failed %v", err)
	}

	fmt.Printf("password hash %v", pass)

	newUser := models.User{
		ID:           1,
		Name:         null.StringFrom(username),
		Username:     null.StringFrom(username),
		PasswordHash: null.StringFrom(pass),
	}

	newUser.Insert(ctx, db, boil.Infer())
}
