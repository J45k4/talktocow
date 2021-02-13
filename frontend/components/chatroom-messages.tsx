import { useEffect, useState } from "react"
import { ChatroomMessage, subscribeToNewChatroomMessages, unsubscribeToNewChatroomMessages } from "../logic/chatroom-messages"
import { OrderedMap } from "immutable"


export const useChatroomMessages = (chatroomId) => {
    const [currentChatroomMessages, setChatroomMessages] = useState(OrderedMap<string, ChatroomMessage>())

    useEffect(() => {
        let deezNuts = OrderedMap<string, ChatroomMessage>()

        function handle(chatroomMessages: ChatroomMessage[]) {
            // const message = chatroomMessages[0];

            // setChatroomMessages(currentChatroomMessages.set(message.reference, message))

            deezNuts = deezNuts.withMutations(m => {
                for (const chatroomMessage of chatroomMessages) {
                    m.set(chatroomMessage.reference, chatroomMessage)
                }
            }).sort((a, b) => {
                const aWritenAt = new Date(a.writenAt).getTime()
                const bWritenAt = new Date(b.writenAt).getTime()

                console.log("aWritenAt", aWritenAt)
                console.log("bWritenAt", bWritenAt)

                return aWritenAt - bWritenAt
            })
            
            console.log("deeznuts", deezNuts)

            setChatroomMessages(deezNuts)
        }

        subscribeToNewChatroomMessages(chatroomId, handle)

        return () => {
            unsubscribeToNewChatroomMessages(chatroomId, handle)
        }
    }, [])

    // console.log("useChatroomMessages currentChatroomMessages", currentChatroomMessages)

    return currentChatroomMessages.reverse().values()
}