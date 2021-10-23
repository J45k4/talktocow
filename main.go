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

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Sql open error ", err)

		return
	}

	chatroomEventbus := chatroom.NewChatroomEventbus()

	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"http://localhost:3080"}
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AddAllowHeaders("authorization")
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

	r.Run(":12001")

	// app := iris.Default()

	// signer := jwt.NewSigner(jwt.HS256, loadPrivateKey(), 12*time.Hour)
	// verifier := jwt.NewVerifier(jwt.HS256, loadPrivateKey())
	// verifyMiddleware := verifier.Verify(func() interface{} {
	// 	return new(UserSession)
	// })
	// app.UseRouter()

	// mvc.Configure(app.Party(""))

	// app.Post("/api/login", func(ctx iris.Context) {
	// 	p := LoginPayload{}

	// 	ctx.ReadJSON(&p)

	// 	fmt.Printf("New login attempt {%v}", p)

	// 	var user models.User

	// 	err := models.Users(
	// 		qm.Where("username = ?", p.Username),
	// 	).Bind(ctx.Request().Context(), db, &user)

	// 	if err != nil {
	// 		ctx.StatusCode(iris.StatusForbidden)
	// 		ctx.EndRequest()

	// 		return
	// 	}

	// 	if auth.CheckPasswordHash(p.Password, user.PasswordHash.String) == false {
	// 		ctx.StatusCode(iris.StatusForbidden)
	// 		ctx.EndRequest()

	// 		return
	// 	}

	// 	token, err := signer.Sign(UserSession{
	// 		UserID: int32(user.ID),
	// 		Name:   user.Name.String,
	// 	})

	// 	if err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		ctx.EndRequest()

	// 		return
	// 	}

	// 	resp := LoginResponse{
	// 		Token: string(token),
	// 	}

	// 	ctx.JSON(resp)
	// })

	// app.Use(verifyMiddleware)

	// app.Get("/api/me", func(ctx iris.Context) {

	// })

	// //func(w http.ResponseWriter, r *http.Request)

	// connections := make(map[*websocket.Conn]bool)

	// app.Get("/api/socket", func(ctx iris.Context) {
	// 	w := ctx.ResponseWriter()
	// 	r := ctx.Request()

	// 	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	// 	if _, ok := err.(websocket.HandshakeError); ok {
	// 		http.Error(w, "Not websocket handshake", 400)
	// 	} else if err != nil {
	// 		return
	// 	}

	// 	userSession := jwt.Get(ctx).(*UserSession)

	// 	fmt.Printf("New sokcet session %v\n", userSession)

	// 	connections[ws] = true

	// 	go func() {
	// 		for {
	// 			msg := WebsocketReceiveMessage{}

	// 			err = ws.ReadJSON(&msg)

	// 			if err != nil {
	// 				delete(connections, ws)

	// 				return
	// 			}

	// 			fmt.Printf("new message %v\n", msg)

	// 			if msg.MessageToChatRoom != nil {
	// 				transmittedAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatRoom.TransmitTime)
	// 				//createdAt, _ := time.Parse(time.RFC3339Nano, msg.MessageToChatRoom.CreateTime)

	// 				newMessage := models.Message{
	// 					MessageText:      null.NewString(msg.MessageToChatRoom.MessageText, true),
	// 					ServerReceivedAt: time.Now(),
	// 					UserID:           int(userSession.UserID),
	// 					Platform:         null.StringFrom("talktocow"),
	// 					ChatroomID:       1,
	// 					TransmitedAt:     transmittedAt,
	// 				}

	// 				messageInserErr := newMessage.Insert(context.Background(), db, boil.Infer())

	// 				if messageInserErr != nil {
	// 					fmt.Println("Message insert failed", messageInserErr)
	// 				}

	// 				newChatroomMessage := NewChatroomMessage{
	// 					MessageText:   msg.MessageToChatRoom.MessageText,
	// 					FromUserName:  userSession.Name,
	// 					TransmittedAt: msg.MessageToChatRoom.TransmitTime,
	// 				}

	// 				transmitMessage := WebsocketTransmitMessage{
	// 					NewChatroomMessage: &newChatroomMessage,
	// 				}

	// 				fmt.Printf("Sending message %v to %v", transmitMessage, userSession)

	// 				for c, _ := range connections {
	// 					c.WriteJSON(transmitMessage)
	// 				}
	// 			}

	// 			// for c, _ := range connections {
	// 			// 	c.WriteJSON(WebsocketSendMessage{
	// 			// 		MessageText: msg.MessageText,
	// 			// 	})
	// 			// }
	// 		}
	// 	}()
	// })

	// app.Get("/api/messages", func(ctx iris.Context) {
	// 	rows := []MessageAndUser{}

	// 	err := models.NewQuery(
	// 		qm.Select("messages.*", "users.*"),
	// 		qm.OrderBy("transmited_at desc"),
	// 		qm.Limit(35),
	// 		qm.From("messages"),
	// 		qm.InnerJoin("users on messages.user_id = users.id"),
	// 	).Bind(ctx.Request().Context(), db, &rows)

	// 	if err != nil {
	// 		fmt.Println("Messages fetch failed", err)

	// 		ctx.SetErr(err)
	// 	}

	// 	for _, row := range rows {
	// 		row.PasswordHash = null.StringFrom("")
	// 	}

	// 	ctx.JSON(rows)
	// })

	// app.Get("/api/chatroom/{chatroom:int32}/messages", func(ctx iris.Context) {

	// })

	// app.Get("/api/chatroom", func(ctx iris.Context) {
	// 	fmt.Println("helo juuri")

	// 	ctx.Writef("Hello world")
	// })

	// app.Use(func(ctx iris.Context) {
	// 	fmt.Println("lol middleware")

	// 	ctx.StatusCode(iris.StatusUnauthorized)
	// 	ctx.EndRequest()
	// })

	// app.Listen(":12001")
}
