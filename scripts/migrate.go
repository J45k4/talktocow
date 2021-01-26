package main

import (
	"database/sql"
	"fmt"

	"github.com/j45k4/talktocow/config"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	connectionString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", config.DBName, config.DBUser, config.DBPassword, config.DbHost, config.DBPort)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Sql open error ", err)

		panic("failed to open database")
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)

	if err != nil {
		fmt.Printf("err %v \n", err)

		panic("Failed to execute migrations")
	}
}
