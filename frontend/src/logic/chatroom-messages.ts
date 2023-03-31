import { getSession } from "./session-manager";
import { MessageFromServer } from "./websocket-types";
import { sendMessageToServer, subscribeToNewMessages, subscribeToSocketStatusChanged } from "./websocket-conn";
import { v4 } from "uuid"
import { getJson } from "../utility/talktocow-api-helpers";
import { ChatroomMessage } from "../types";

type SubscriberCallback = (payload: ChatroomMessage[]) => void;

const subscribers = new Map<string, Set<SubscriberCallback>>()

const chatroomMessagesByChatroom = new Map<string, Map<string, ChatroomMessage>>();

const notifyChatroomSubscribers = (chatroomId: string, chatroomMessages: ChatroomMessage[]) => {
    let chatroomSubscribers = subscribers.get(chatroomId)

    for (const [, sub] of chatroomSubscribers.entries()) {
        sub(chatroomMessages)
    }
}

const provideChatroomMessages = (chatroomId: string) => {
    let chatroomMessages = chatroomMessagesByChatroom.get(chatroomId)

    if (!chatroomMessages) {
        chatroomMessages = new Map()
        chatroomMessagesByChatroom.set(chatroomId, chatroomMessages)
    }

    return chatroomMessages
}

function handle(message: MessageFromServer) {
    console.log("message forom server", message)

    if (message.newChatroomMessage) {
        const newChatroomMessage = message.newChatroomMessage

        const chatroomMessages = provideChatroomMessages(newChatroomMessage.chatroomId)

        const existingMessage = chatroomMessages.get(newChatroomMessage.reference)

        if (existingMessage) {
            existingMessage.serverReceivedAt = newChatroomMessage.serverReceivedAt

            notifyChatroomSubscribers(newChatroomMessage.chatroomId, [existingMessage])
        } else {
            const chatroomMessage: ChatroomMessage = {
                messageText: newChatroomMessage.messageText,
                userId: newChatroomMessage.userId,
                userName: newChatroomMessage.userName,
                writenAt: newChatroomMessage.writenAt,
                messageId: newChatroomMessage.messageId,
                platform: newChatroomMessage.platform,
                reference: newChatroomMessage.reference,
                serverReceivedAt: newChatroomMessage.serverReceivedAt,
                transmitedAt: newChatroomMessage.transmitedAt
            }

            chatroomMessages.set(newChatroomMessage.reference, chatroomMessage)

            notifyChatroomSubscribers(newChatroomMessage.chatroomId, [chatroomMessage])
        }
    }

    console.log("Messages by chatroom:", chatroomMessagesByChatroom)
}

subscribeToNewMessages(handle)

export const getChatroomEvents = (chatroomId: string) => {
    const chatroomEvents = chatroomMessagesByChatroom.get(chatroomId)

    if (!chatroomEvents) {
        return []
    }

    return chatroomEvents
}

export const sendMessageToChatroom = (chatroomId: string, message: string) => {
    const session = getSession()

    const writenAt = new Date().toISOString()

    const chatroomMessage: ChatroomMessage = {
        userId: session.userId,
        userName: session.username,
        messageText: message,
        writenAt: writenAt,
        reference: v4()
    }

    const chatroomMessages = provideChatroomMessages(chatroomId)

    chatroomMessages.set(chatroomMessage.reference, chatroomMessage)

    notifyChatroomSubscribers(chatroomId, [chatroomMessage])

    const sendSuccess = sendMessageToServer({
        messageToChatroom: {
            chatroomId: chatroomId,
            messageText: message,
            reference: chatroomMessage.reference,
            transmitedAt: new Date().toISOString(),
            writenAt: writenAt
        }
    })

    if (sendSuccess == false) {
        console.log("send did not success lol")
    }
}

export const subscribeToNewChatroomMessages = (chatroomId: string, cb: SubscriberCallback) => {
    let chatroomSubscribers = subscribers.get(chatroomId)

    if (!chatroomSubscribers) {
        chatroomSubscribers = new Set()
        subscribers.set(chatroomId, chatroomSubscribers)
    }

    chatroomSubscribers.add(cb)
}

export const unsubscribeToNewChatroomMessages = (chatroomId: string, cb: SubscriberCallback) => {
    if (subscribers.has(chatroomId)) {
        return
    }

    const chatroomSubscribers = subscribers.get(chatroomId)
    chatroomSubscribers.delete(cb)
}

export const fetchChatroomMessages = async (chatroomId: string) => {
    const res = await getJson<ChatroomMessage[]>(`/api/chatroom/${chatroomId}/messages`)

    if (res.payload) {
        const chatroomMessages = provideChatroomMessages(chatroomId)

        for (const msg of res.payload) {
            chatroomMessages.set(msg.reference, msg)
        }

        notifyChatroomSubscribers(chatroomId, res.payload)
    }
}

subscribeToSocketStatusChanged(s => {

})