//go:generate sqlboiler psql

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/j45k4/talktocow/bot"
	"github.com/j45k4/talktocow/chatroom"
	"github.com/j45k4/talktocow/config"
	"github.com/j45k4/talktocow/eventbus"
	"github.com/j45k4/talktocow/routes"
	_ "github.com/lib/pq"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	migrate "github.com/rubenv/sql-migrate"
)

func corsOriginAllowed(origin string) bool {
	for _, allowedOrigin := range config.CORSAllowOrigins {
		if origin == allowedOrigin {
			return true
		}
	}

	parsedOrigin, err := url.Parse(origin)
	if err != nil {
		return false
	}

	hostname := parsedOrigin.Hostname()
	if hostname == "localhost" || hostname == "puppe" || strings.HasSuffix(hostname, ".local") {
		return true
	}

	ip := net.ParseIP(hostname)
	if ip == nil {
		return false
	}

	return ip.IsLoopback() || ip.IsPrivate()
}

func registerFrontendRoutes(r *gin.Engine, distPath string) {
	if distPath == "" {
		return
	}

	indexPath := filepath.Join(distPath, "index.html")
	if _, err := os.Stat(indexPath); err != nil {
		log.Printf("frontend dist not available at %s", distPath)
		return
	}

	r.NoRoute(func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/api" || strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
			ctx.JSON(http.StatusNotFound, routes.CreateErrorResponse(routes.NotFound, "Not found"))
			return
		}

		requestPath := path.Clean("/" + ctx.Request.URL.Path)
		filePath := filepath.Join(distPath, filepath.FromSlash(strings.TrimPrefix(requestPath, "/")))
		relativePath, err := filepath.Rel(distPath, filePath)
		if err != nil || relativePath == ".." || strings.HasPrefix(relativePath, ".."+string(filepath.Separator)) {
			ctx.JSON(http.StatusNotFound, routes.CreateErrorResponse(routes.NotFound, "Not found"))
			return
		}

		fileInfo, err := os.Stat(filePath)
		if err == nil && !fileInfo.IsDir() {
			ctx.File(filePath)
			return
		}

		ctx.File(indexPath)
	})
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

	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{}
	corsConfig.AllowOriginFunc = corsOriginAllowed
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
	r.POST("/api/token", routes.HandleTokenLogin)
	r.POST("/api/passkeys/login/begin", routes.HandlePasskeyLoginBegin)
	r.POST("/api/passkeys/login/finish", routes.HandlePasskeyLoginFinish)
	r.POST("/api/logout", routes.HandleLogout)
	r.GET("/api/ws", routes.HandleWs)

	registerFrontendRoutes(r, config.FrontendDistPath)

	authenticated := r.Group("", routes.SessionMiddleware)

	authenticated.GET("/api/passkeys", routes.GetPasskeys)
	authenticated.POST("/api/passkeys/registration/begin", routes.HandlePasskeyRegistrationBegin)
	authenticated.POST("/api/passkeys/registration/finish", routes.HandlePasskeyRegistrationFinish)
	authenticated.GET("/api/users", routes.GetUsers)
	authenticated.GET("/api/chatrooms", routes.GetChatrooms)
	authenticated.POST("/api/chatroom", routes.CreateChatroom)
	authenticated.GET("/api/chatroom/:chatroomId", routes.GetChatroom)
	authenticated.PATCH("/api/chatroom/:chatroomId", routes.PatchChatroom)
	authenticated.POST("/api/chatroom/:chatroomId/member", routes.AddChatroomMember)
	authenticated.DELETE("/api/chatroom/:chatroomId/member/:userId", routes.RemoveChatroomMember)
	authenticated.GET("/api/chatroom/:chatroomId/members", routes.GetChatroomMembers)
	authenticated.GET("/api/chatroom/:chatroomId/messages", routes.GetChatroomMessages)
	authenticated.GET("/api/mychatrooms", routes.GetMyChatrooms)
	authenticated.GET("/api/messages", routes.HandleGetMessages)
	authenticated.GET("/api/socket", routes.HandleSocket)
	authenticated.POST("/api/files", routes.UploadFile)
	authenticated.GET("/api/files/:fileId", routes.GetFile)
	authenticated.DELETE("/api/files/:fileId", routes.DeleteFile)
	authenticated.POST("/api/diary/entry", routes.CreateDiaryEntry)
	authenticated.GET("/api/diary/entries", routes.GetDiaryEntries)
	authenticated.GET("/api/diary/entries/count", routes.GetDiaryEntriesCount)
	authenticated.GET("/api/diary/labels", routes.GetDiaryLabels)
	authenticated.GET("/api/diary/entry/:diaryEntryId", routes.GetDiaryEntry)
	authenticated.PUT("/api/diary/entry/:diaryEntryId", routes.UpdateDiaryEntry)
	authenticated.DELETE("/api/diary/entry/:diaryEntryId", routes.DeleteDiaryEntry)
	authenticated.GET("/api/diary/entry/:diaryEntryId/pictures", routes.GetDiaryEntryPictures)
	authenticated.POST("/api/diary/entry/:diaryEntryId/picture", routes.UploadDiaryEntryPicture)
	authenticated.GET("/api/diary/entry/:diaryEntryId/picture/:pictureId", routes.GetDiaryEntryPicture)
	authenticated.DELETE("/api/diary/entry/:diaryEntryId/picture/:pictureId", routes.DeleteDiaryEntryPicture)

	authenticated.POST("/api/diary/entry/:diaryEntryId/comment", routes.CreateDiaryEntryComment)
	authenticated.POST("/api/diary/entry/:diaryEntryId/comment/:commentId", routes.UpdateDiaryEntryComment)
	authenticated.DELETE("/api/diary/entry/:diaryEntryId/comment/:commentId", routes.DeleteDiaryEntryComment)
	authenticated.GET("/api/diary/entry/:diaryEntryId/comments", routes.GetDiaryEntryComments)
	authenticated.GET("/api/diary/entry/:diaryEntryId/comments/count", routes.GetDiaryEntryCommentsCount)

	authenticated.POST("/api/pushovertoken", routes.CreatePushoverToken)
	authenticated.GET("/api/pushovertokens", routes.GetPushoverTokens)
	authenticated.DELETE("/api/pushovertoken/:pushoverTokenId", routes.DeletePushoverToken)
	authenticated.GET("/api/pushovertoken/:pushoverTokenId", routes.GetPushoverToken)

	authenticated.GET("/api/course/:courseId", routes.GetCourse)
	authenticated.GET("/api/courses", routes.GetCourses)
	authenticated.POST("/api/course", routes.CreateCourse)
	authenticated.PUT("/api/course/:courseId", routes.UpdateCourse)
	authenticated.GET("/api/course/:courseId/mymeta", routes.GetCourseMeta)

	authenticated.GET("/api/course/:courseId/homeworks", routes.GetHomeworks)
	authenticated.POST("/api/course/:courseId/homework", routes.CreateHomework)
	authenticated.PUT("/api/course/:courseId/homework/:homeworkId", routes.UpdateHomework)
	authenticated.GET("/api/course/:courseId/homework/:homeworkId", routes.GetHomework)
	authenticated.POST("/api/course/:courseId/homework/:homeworkId/submit", routes.SubmitHomework)

	authenticated.POST("/api/course/:courseId/homework/:homeworkId/submission", routes.SubmitHomework)
	authenticated.GET("/api/course/:courseId/homework/:homeworkId/submissions", routes.GetHomeworkSubmissions)
	authenticated.GET("/api/course/:courseId/student/:userId/submissions", routes.GetStudentSubmissions)
	authenticated.GET("/api/course/:courseId/students", routes.GetCourseStudents)
	authenticated.POST("/api/course/:courseId/student", routes.AddUserToCourse)

	r.Run(":12001")
}
