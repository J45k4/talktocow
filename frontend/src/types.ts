
export type SendMessage = {
	type: "sendMessage"
	chatroomId: number
	message: string
}

export type Authenticate = {
	type: "authenticate"
	token: string
}

export type AskQuestion = {
	type: "askQuestion"
}

export type MessageToServer = SendMessage |
	Authenticate |
	AskQuestion