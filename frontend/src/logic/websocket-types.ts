// Message from server

export interface UserStatus {
    online: boolean
    userId: number
    username: string
    lastSeen: string
    sleeping: boolean
    timestamp: string
}

export interface NewChatroomMessage {
    chatroomId: string
    userId: string
    userName: string
    messageId: string
    messageText: string    
    writtenAt: string
    transmitedAt: string
    serverReceivedAt?: string
    platform: string
    reference: string
}

export interface MessageFromServer {
    type: "messageFromServer"
    changedUserStatus?: UserStatus
    newChatroomMessage?: NewChatroomMessage
}

// Message to server

export interface MessageToChatroom {
    messageText: string
    chatroomId: string
    writtenAt: string
    transmitedAt: string
    reference: string
}

export interface MessageToServer {
    messageToChatroom?: MessageToChatroom
    iamHere?: boolean
}

export type Connected = {
    type: "connected"
}

export type Disconnected = {
    type: "disconnected"
}

export type ConnEvent = MessageFromServer | 
    Connected | 
    Disconnected