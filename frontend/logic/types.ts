// Messages from server

export interface Message {
    id: number
    messageText: string
    writenAt: string
    transmitedAt: string
    platform: string
    chatroomId: string
}

export interface ChatroomEvent {
    id: number
    chatroomId: number
    message: Message
    createdAt: string
}

export interface UserStatus {
    online: boolean
    userId: number
    username: string
    lastSeen: string
    sleeping: boolean
    timestamp: string
}

export interface MessageFromServer {
    newChatroomEvent?: ChatroomEvent
    changedUserStatus?: UserStatus
}

// Message to server

export interface MessageToChatroom {
    messageText: string
    chatroomId: string
    writedAt: string
    transmitedAt: string
}

export interface MessageToServer {
    messageToChatroom?: MessageToChatroom
    iamHere?: boolean
}