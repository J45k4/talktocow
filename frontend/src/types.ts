
export type SendMessage = {
	type: "sendMessage"
} & ChatroomMessage

export type Authenticate = {
	type: "authenticate"
	token: string
}

export type AskQuestion = {
	type: "askQuestion"
}

export type MessageToServer = (SendMessage |
	Authenticate |
	AskQuestion) & {
		transmitedAt?: string
	}

export type ChatroomMessage = {
	userId: string
	userName: string
	messageId?: string
	messageText: string
	writtenAt: string
	transmitedAt?: string
	serverReceivedAt?: string
	platform?: string
	reference?: string
	chatroomId: string
}

export type ChatroomMessages = {
	type: "chatroomMessages"
	messages: ChatroomMessage[]
}

export type MessageFromServer = ChatroomMessages

export type Chatroom = {
	id: string
	name: string
}

export type User = {
	id: string
	name: string
}