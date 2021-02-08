import { useEffect, useState } from "react"
import { ChatroomEvent, getChatroomEvents, subscribeToChatroomEvents, unsubscribeToChatroomEvents } from "../logic/chatroom-manager"

export const useChatroomEvents = (chatroomId: string) => {
    const [events, setEvents] = useState(getChatroomEvents(chatroomId))

    useEffect(() => {
        function handle(events: ChatroomEvent[]) {
            setEvents(events)
        }

        subscribeToChatroomEvents(chatroomId, handle)

        return () => {
            unsubscribeToChatroomEvents(chatroomId, handle)
        }
    }, [])

    return events
}