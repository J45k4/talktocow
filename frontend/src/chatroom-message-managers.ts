import { v4 } from "uuid"
import { chatroomMessageStore } from "./chatroom-message-store"
import { eventbus } from "./eventbus"
import { createLogger } from "./logger"
import { getSession } from "./logic/session-manager"
import { ChatroomMessage } from "./types"
import { getJson } from "./utility/talktocow-api-helpers"
import { ws } from "./ws"

const logger = createLogger("chatroomMessageManager")

export const sendMessageToChatroom = (chatroomId: string, message: string) => {
	const session = getSession()

	const writenAt = new Date().toISOString()

    const chatroomMessage: ChatroomMessage = {
        userId: session.userId,
        userName: session.username,
        messageText: message,
        writenAt: writenAt,
        reference: v4(),
		chatroomId: chatroomId
    }

	chatroomMessageStore.addMessage(chatroomMessage)

	eventbus.publish("chatroomMessages", [chatroomMessage])

	ws.send({
		type: "sendMessage",
		...chatroomMessage
	})
}

export const loadChatroomMessages = (chatroomId: string, n: number) => {
	getJson<ChatroomMessage[]>(`/api/chatroom/${chatroomId}/messages`)
		.then(res => {
			if (res.payload) {
				logger.info("Loaded chatroom messages", res.payload)

				chatroomMessageStore.addMessages(res.payload)
				eventbus.publish("chatroomMessages", res.payload)
			}
		})
}