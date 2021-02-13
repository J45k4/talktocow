import { MessageFromServer } from "./websocket-types";
import {subscribeToNewMessages } from "./websocket-conn";


export interface UserStatus {
    online: boolean
    userId: number
    username: string
    lastSeen: string
    sleeping: boolean
}

const userStatus = new Map<string, UserStatus>();


export const getUserStatus = (userId: string) => {
    return userStatus.get(userId)
}

function handle(message: MessageFromServer) {

}

subscribeToNewMessages(handle)
