package routes

import "github.com/volatiletech/null/v8"

type ChatroomMessage struct {
	ChatroomID   string      `json:"chatroomId" boil:"chatroom_id"`
	UserID       string      `json:"userId" boil:"user_id"`
	UserName     string      `json:"userName" boil:"user_name"`
	MessageID    string      `json:"messageId" boil:"message_id"`
	MessageText  string      `json:"messageText" boil:"message_text"`
	WritenAt     string      `json:"writenAt" boil:"writen_at"`
	TransmitedAt string      `json:"transmitedAt" boil:"transmited_at"`
	Platform     string      `json:"platform" boil:"platform"`
	Reference    null.String `json:"reference" boil:"reference"`
}

type MessageToChatroom struct {
	MessageText  string `json:"messageText"`
	ChatroomID   string `json:"chatroomId"`
	WritenAt     string `json:"writenAt"`
	TransmitedAt string `json:"transmitedAt"`
	Reference    string `json:"reference"`
}

type UserStatus struct {
	Online    bool   `json:"online"`
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	Lastseen  string `json:"lastseen"`
	Sleeping  bool   `json:"sleeping"`
	Timestamp string `json:"timestamp"`
}

type NewChatroomMessage struct {
	ChatroomID       string `json:"chatroomId"`
	UserID           string `json:"userId"`
	UserName         string `json:"userName"`
	MessageID        string `json:"messageId"`
	MessageText      string `json:"messageText"`
	WritenAt         string `json:"writenAt"`
	TransmitedAt     string `json:"transmitedAt"`
	ServerReceivedAt string `json:"serverReceivedAt"`
	Platform         string `json:"platform"`
	Reference        string `json:"reference"`
}
