//go:generate sqlboiler psql

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/j45k4/talktocow/bot"
	"github.com/j45k4/talktocow/chatroom"
	"github.com/j45k4/talktocow/config"
	"github.com/j45k4/talktocow/eventbus"
	_ "github.com/lib/pq"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	log.Printf("Private key path %v", config.PrivateKeyPath)
	log.Printf("Public key path %v", config.PublicKeyPath)

	connectionString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", config.DBName, config.DBUser, config.DBPassword, config.DbHost, config.DBPort)

	log.Printf("connection string %v", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("opening database connection failed %v", err)
		return
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Printf("failed to execute migrations %v", err)
		panic("Failed to execute migrations")
	}

	bot.InitializeBots(db)

	eventbus := eventbus.New()
	cowgpt := bot.CowGPT{
		Eventbus:   eventbus,
		Client:     openai.NewClient(option.WithAPIKey(config.OpenAIApiKey)),
		CowGPTUser: bot.GetCowGPTUser(db),
		Ctx:        context.Background(),
		DB:         db,
	}
	go cowgpt.Run()

	chatroomEventbus := chatroom.NewChatroomEventbus()
	r := setupRouter(db, eventbus, chatroomEventbus)
	registerFrontendRoutes(r, config.FrontendDistPath)

	r.Run(":12001")
}
