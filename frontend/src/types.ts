
export type SendMessage = {
	type: "sendMessage"
} & ChatroomMessage

export type Authenticate = {
	type: "authenticate"
	token: string
	deivceId: string
}

export type AskQuestion = {
	type: "askQuestion"
}

export type WebRTCOffer = {
	type: "createWebRTCOffer"
	sdp: string
	userId: string
	deviceId: string
}

export type WebRTCOfferAnswer = {
	type: "createWebRTCOfferAnswer"
	sdp: string
	userId: string
	deviceId: string
}

export type CreateCall = {
	type: "createCall"
	chatroomId?: string
	userIds?: string[]
}

export type JoinCall = {
	type: "joinCall"
	callId: string
}

export type InviteUserToCall = {
	type: "inviteUserToCall"
	callId: string
	userId: string
}

export type AcceptUserToCall = {
	type: "acceptUserToCall"
	callId: string
	userId: string
}

export type RejectUserToCall = {
	type: "rejectUserToCall"
	callId: string
	userId: string
}

export type LeaveCall = {
	type: "leaveCall"
	callId: string
}

export type MessageToServer = (
	SendMessage |
	Authenticate |
	WebRTCOffer |
	WebRTCOfferAnswer |
	AskQuestion |
	CreateCall |
	JoinCall |
	InviteUserToCall |
	AcceptUserToCall |
	RejectUserToCall |
	LeaveCall
) & {
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

export type VideoOfferAnswer = {
	type: "videoOfferAnswer"
	sdp: string
}

export type NewIceCandidate = {
	type: "newIceCandidate"
	candidate: string
}

export type UserJoinedCall = {
	type: "userJoinedCall"
	callId: string
	userId: string
}

export type UserLeftCall = {
	type: "userLeftCall"
	callId: string
	userId: string
}

export type UserInvidedToJoinCall = {
	type: "userInvidedToJoinCall"
	callId: string
	userId: string
}

export type UserRequestedToJoinCall = {
	type: "userRequestedToJoinCall"
	callId: string	
	userId: string
}

export type UserWasRejectedToJoinCall = {
	type: "userWasRejectedToJoinCall"
	callId: string
	userId: string
}

export type DeviceJoinedCall = {
	type: "deviceJoinedCall"
	callId: string
	userId: string
	deviceId: string
}

export type DeviceLeftCall = {
	type: "deviceLeftCall"
	callId: string
	userId: string
	deviceId: string
}

export type CallDevices = {
	type: "callDevices"
	callId: string
	devices: {
		userId: string
		deviceId: string
	}[]
}

export type MessageFromServer = ChatroomMessages | 
	VideoOfferAnswer |
	NewIceCandidate |
	UserJoinedCall |
	UserLeftCall |
	UserRequestedToJoinCall |
	UserInvidedToJoinCall |
	UserWasRejectedToJoinCall |
	DeviceJoinedCall |
	DeviceLeftCall |
	CallDevices

export type Chatroom = {
	id: string
	name: string
}

export type User = {
	id: string
	name: string
}