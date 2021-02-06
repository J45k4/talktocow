import { MessageFromServer } from "./types";
import { subscribeToNewMessages } from "./websocket-conn";

export type DayChanged = {
    type: "dayChanged"
    date: string
}

export type NewMessage = {
    type: "newMessage"
    userId: number
    userName: string
    message: string
    writenAt: string
    transmittedAt: string
}

export type ChatroomEvent = DayChanged | NewMessage

const chatroomEventsByChatroom = new Map<number, ChatroomEvent[]>();

function handle(message: MessageFromServer) {
    if (message.newChatroomMessage) {
        const newChatroomMessage = message.newChatroomMessage

        let chatroomEvents = chatroomEventsByChatroom.get(message.newChatroomMessage.chatroomId)

        if (!chatroomEvents) {
            chatroomEvents = []
        }

        chatroomEvents.push({
            type: "newMessage",
            userId: newChatroomMessage.userId,
            userName: newChatroomMessage.userName,
            message: newChatroomMessage.messageText,
            writenAt: newChatroomMessage.writenAt,
            transmittedAt: newChatroomMessage.transmittedAt
        })
    }
}

subscribeToNewMessages(handle)

export const getChatroomEvents = (chatroomId: number) => {
    const chatroomEvents = chatroomEventsByChatroom.get(chatroomId)

    if (!chatroomEvents) {
        return []
    }

    return chatroomEvents
}

export const sendMessageToChatroom = (chatroomId: string, message: string) => {
    
}