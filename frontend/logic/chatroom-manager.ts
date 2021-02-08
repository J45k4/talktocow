import { getSession } from "./session-manager";
import { MessageFromServer } from "./types";
import { sendMessageToServer, subscribeToNewMessages, subscribeToSocketStatusChanged } from "./websocket-conn";

export type NewMessage = {
    type: "newMessage"
    userId: string
    userName: string
    messageText: string
    writenAt: string
    transmittedAt?: string
    serverReceivedAt?: string
}

export type ChatroomEvent = NewMessage

type SubscriberCallback = (payload: ChatroomEvent[]) => void;

const subscribers = new Map<string, Set<SubscriberCallback>>()

const chatroomEventsByChatroom = new Map<string, ChatroomEvent[]>();

const notifyChatroomSubscribers = (chatroomId: string, events: ChatroomEvent[]) => {
    let chatroomSubscribers = subscribers.get(chatroomId)

    for (const [, sub] of chatroomSubscribers.entries()) {
        sub(events)
    }
}

function handle(message: MessageFromServer) {
    // if (message.newChatroomMessage) {
    //     const newChatroomMessage = message.newChatroomMessage

    //     let chatroomEvents = chatroomEventsByChatroom.get(message.newChatroomMessage.chatroomId)

    //     if (!chatroomEvents) {
    //         chatroomEvents = []
    //     }

    //     chatroomEvents.push({
    //         type: "newMessage",
    //         userId: newChatroomMessage.userId,
    //         userName: newChatroomMessage.userName,
    //         message: newChatroomMessage.messageText,
    //         writenAt: newChatroomMessage.writenAt,
    //         transmittedAt: newChatroomMessage.transmittedAt
    //     })
    // }
}

subscribeToNewMessages(handle)

export const getChatroomEvents = (chatroomId: string) => {
    const chatroomEvents = chatroomEventsByChatroom.get(chatroomId)

    if (!chatroomEvents) {
        return []
    }

    return chatroomEvents
}

export const sendMessageToChatroom = (chatroomId: string, message: string) => {
    const events = getChatroomEvents(chatroomId);

    const session = getSession()

    const writenAt = new Date().toISOString()

    const newMessage: NewMessage = {
        type: "newMessage",
        messageText: message,
        writenAt: writenAt,
        userId: session.userId,
        userName: session.username
    }

    const newEvents = [
        newMessage,
        ...events        
    ]

    chatroomEventsByChatroom.set(chatroomId, newEvents)

    sendMessageToServer({
        messageToChatroom: {
            chatroomId: chatroomId,
            messageText: message,
            transmitedAt: new Date().toISOString(),
            writedAt: writenAt
        }
    })

    notifyChatroomSubscribers(chatroomId, newEvents)
}

export const subscribeToChatroomEvents = (chatroomId: string, cb: SubscriberCallback) => {
    let chatroomSubscribers = subscribers.get(chatroomId)

    if (!chatroomSubscribers) {
        chatroomSubscribers = new Set()
        subscribers.set(chatroomId, chatroomSubscribers)
    }

    chatroomSubscribers.add(cb)
}

export const unsubscribeToChatroomEvents = (chatroomId: string, cb: SubscriberCallback) => {
    if (subscribers.has(chatroomId)) {
        return
    }

    const chatroomSubscribers = subscribers.get(chatroomId)
    chatroomSubscribers.delete(cb)
}

subscribeToSocketStatusChanged(s => {

})