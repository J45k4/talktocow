import React, { useEffect, useState } from "react"
import { getSession, SessionChangeNotify, subscribeToSessionEvents, unsubscribeToSessionEvents } from "../logic/session-manager"
import { MessageFromServer } from "../logic/websocket-types"
import { getUserStatus } from "../logic/user-manager"
import {subscribeToNewMessages, unsubscribeToMessages } from "../logic/websocket-conn"


export const useIsUserStatus = (userId: string) => {
    const [userStatus, setUserStatus] = useState(getUserStatus(userId))

    useEffect(() => {
        function handle(newMessage: MessageFromServer) {

        }

        subscribeToNewMessages(handle)

        return () => {
            unsubscribeToMessages(handle)
        }
    }, [])

    return userStatus
}

export const useIsUserOnline = (userId: string) => {
    const userStatus = useIsUserStatus(userId)

    return userStatus?.online ? true : false
}

export const useLastseen = (userId: string) => {
    const userStatus = useIsUserStatus(userId)

    return userStatus.lastSeen
}

export const useYourUserId = () => {
    const [currentUserId, setCurrentUserId] = useState(getSession().userId)

    useEffect(() => {
        function handle(u: SessionChangeNotify) {
            setCurrentUserId(u.userId)
        }

        subscribeToSessionEvents(handle)

        return () => {
            unsubscribeToSessionEvents(handle)
        }
    }, [])


    return currentUserId
}