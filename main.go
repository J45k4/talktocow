//go:generate sqlboiler psql

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/chatroom"
	"github.com/j45k4/talktocow/config"
	"github.com/j45k4/talktocow/models"
	"github.com/j45k4/talktocow/routes"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password`
}

type UserSession struct {
	UserID int32
	Name   string
}

type LoginResponse struct {
	Token        string `json:"token"`
	ErrorMessage string `json:"errorMessage"`
}

func loadPrivateKey() []byte {
	privateKeyBytes, err := ioutil.ReadFile(config.PrivateKeyPath)

	if err != nil {
		panic("Reading private key failed")
	}

	return privateKeyBytes
}

func loadPublicKey() []byte {
	publicKeyBytes, err := ioutil.ReadFile(config.PublicKeyPath)

	if err != nil {
		panic("No public key found")
	}

	return publicKeyBytes
}

//func EchoServer(ws *websocket.Conn) {
//	io.Copy(ws, ws)
//}

type MessageToChatroom struct {
	MessageText  string `json:"messageText"`
	CreateTime   string `json:"createTime"`
	TransmitTime string `json:"transmitTime"`
}

type WebsocketReceiveMessage struct {
	MessageToChatRoom *MessageToChatroom `json:"messageToChatroom"`
}

type NewChatroomMessage struct {
	MessageText   string `json:"messageText"`
	FromUserName  string `json:"fromUserName"`
	TransmittedAt string `json:"transmittedAt"`
}

type WebsocketTransmitMessage struct {
	NewChatroomMessage *NewChatroomMessage `json:"newChatroomMessage"`
}

type MessageAndUser struct {
	models.Message `boil:",bind"`
	models.User    `boil:",bind"`
}

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

	chatroomEventbus := chatroom.NewChatroomEventbus()

	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowCredentials = true
	// corsConfig.AllowOrigins = []string{"http://localhost:3080", "https://talktocow.dy.fi/"}
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AddAllowHeaders("authorization")
	corsConfig.AddAllowHeaders("x-device-id")
	r.Use(cors.New(corsConfig))

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		chatroom.SetChatroomEventbus(ctx, chatroomEventbus)
	})

	r.POST("/api/login", routes.HandleLogin)

	r.Use(routes.SessionMiddleware)

	r.GET("/api/chatroom/:chatroomId/messages", routes.GetChatroomMessages)
	r.GET("/api/messages", routes.HandleGetMessages)
	r.GET("/api/socket", routes.HandleSocket)
	r.POST("/api/diary/entry", routes.CreateDiaryEntry)
	r.GET("/api/diary/entries", routes.GetDiaryEntries)
	r.GET("/api/diary/entries/count", routes.GetDiaryEntriesCount)
	r.GET("/api/diary/entry/:diaryEntryId", routes.GetDiaryEntry)
	r.PUT("/api/diary/entry/:diaryEntryId", routes.UpdateDiaryEntry)
	r.DELETE("/api/diary/entry/:diaryEntryId", routes.DeleteDiaryEntry)

	r.Run(":12001")

	r.Run(":12001")
}
