//go:generate sqlboiler psql

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/bot"
	"github.com/j45k4/talktocow/chatroom"
	"github.com/j45k4/talktocow/config"
	"github.com/j45k4/talktocow/eventbus"
	"github.com/j45k4/talktocow/routes"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sashabaranov/go-openai"
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
		Client:     openai.NewClient(config.OpenAIApiKey),
		CowGPTUser: bot.GetCowGPTUser(db),
		Ctx:        context.Background(),
		DB:         db,
	}

	go cowgpt.Run()

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
		ctx.Set("eventbus", eventbus)
		chatroom.SetChatroomEventbus(ctx, chatroomEventbus)
	})

	r.POST("/api/login", routes.HandleLogin)
	r.GET("/api/ws", routes.HandleWs)

	r.Use(routes.SessionMiddleware)

	r.GET("/api/users", routes.GetUsers)
	r.GET("/api/chatrooms", routes.GetChatrooms)
	r.POST("/api/chatroom", routes.CreateChatroom)
	r.GET("/api/chatroom/:chatroomId", routes.GetChatroom)
	r.PATCH("/api/chatroom/:chatroomId", routes.PatchChatroom)
	r.GET("/api/chatroom/:chatroomId/members", routes.GetChatroomMembers)
	r.GET("/api/chatroom/:chatroomId/messages", routes.GetChatroomMessages)
	r.GET("/api/mychatrooms", routes.GetMyChatrooms)
	r.GET("/api/messages", routes.HandleGetMessages)
	r.GET("/api/socket", routes.HandleSocket)
	r.POST("/api/diary/entry", routes.CreateDiaryEntry)
	r.GET("/api/diary/entries", routes.GetDiaryEntries)
	r.GET("/api/diary/entries/count", routes.GetDiaryEntriesCount)
	r.GET("/api/diary/entry/:diaryEntryId", routes.GetDiaryEntry)
	r.PUT("/api/diary/entry/:diaryEntryId", routes.UpdateDiaryEntry)
	r.DELETE("/api/diary/entry/:diaryEntryId", routes.DeleteDiaryEntry)

	r.POST("/api/diary/entry/:diaryEntryId/comment", routes.CreateDiaryEntryComment)
	r.POST("/api/diary/entry/:diaryEntryId/comment/:commentId", routes.UpdateDiaryEntryComment)
	r.DELETE("/api/diary/entry/:diaryEntryId/comment/:commentId", routes.DeleteDiaryEntryComment)
	r.GET("/api/diary/entry/:diaryEntryId/comments", routes.GetDiaryEntryComments)
	r.GET("/api/diary/entry/:diaryEntryId/comments/count", routes.GetDiaryEntryCommentsCount)

	r.POST("/api/pushovertoken", routes.CreatePushoverToken)
	r.GET("/api/pushovertokens", routes.GetPushoverTokens)
	r.DELETE("/api/pushovertoken/:pushoverTokenId", routes.DeletePushoverToken)
	r.GET("/api/pushovertoken/:pushoverTokenId", routes.GetPushoverToken)

	r.GET("/api/course/:courseId", routes.GetCourse)
	r.GET("/api/courses", routes.GetCourses)
	r.POST("/api/course", routes.CreateCourse)
	r.PUT("/api/course/:courseId", routes.UpdateCourse)
	r.GET("/api/course/:courseId/mymeta", routes.GetCourseMeta)

	r.GET("/api/course/:courseId/homeworks", routes.GetHomeworks)
	r.POST("/api/course/:courseId/homework", routes.CreateHomework)
	r.PUT("/api/course/:courseId/homework/:homeworkId", routes.UpdateHomework)
	r.GET("/api/course/:courseId/homework/:homeworkId", routes.GetHomework)
	r.POST("/api/course/:courseId/homework/:homeworkId/submit", routes.SubmitHomework)

	r.POST("/api/course/:courseId/homework/:homeworkId/submission", routes.SubmitHomework)
	r.GET("/api/course/:courseId/homework/:homeworkId/submissions", routes.GetHomeworkSubmissions)
	r.GET("/api/course/:courseId/student/:userId/submissions", routes.GetStudentSubmissions)
	r.GET("/api/course/:courseId/students", routes.GetCourseStudents)
	r.POST("/api/course/:courseId/student", routes.AddUserToCourse)

	r.Run(":12001")
}
