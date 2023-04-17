package eventbus

import "time"

type ChatroomMessage struct {
	ID               int       `json:"id"`
	MessageText      string    `json:"messageText"`
	WrittenAt        time.Time `json:"writtenAt"`
	TransmitedAt     time.Time `json:"transmitedAt"`
	ServerReceivedAt time.Time `json:"serverReceivedAt"`
	UserID           int       `json:"userId"`
	ChatroomID       int       `json:"chatroomId"`
	Platform         string    `json:"platform"`
	CreatedAt        time.Time `json:"createdAt"`
	Reference        string    `json:"reference"`
	Bot              bool      `json:"bot"`
}

type VideoOfferEvent struct {
	SDP    string `json:"sdp"`
	UserID int    `json:"userId"`
}

type VideoAnswerEvent struct {
	SDP    string `json:"sdp"`
	UserID int    `json:"userId"`
}

type NewIceCandidateEvent struct {
	UserID    int    `json:"userId"`
	Candidate string `json:"candidate"`
}

type Event struct {
	ChatroomMessage *ChatroomMessage
	VideoOffer      *VideoOfferEvent
	VideoAnswer     *VideoAnswerEvent
	NewIceCandidate *NewIceCandidateEvent
}
