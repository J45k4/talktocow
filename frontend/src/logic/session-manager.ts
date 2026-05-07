import { v4 } from "uuid"

type SubscriberCallback = (payload: SessionChangeNotify) => void;

export interface SessionChangeNotify {
    token: string
    deviceId: string
    username: string
    userId: string
    authMethod?: "password" | "passkey"
}

let token
let deviceId;
let username;
let userId;
let authMethod;

const subscribers = new Set<SubscriberCallback>()

const notifyChanges = () => {
    const notify: SessionChangeNotify = {
        token: token,
        deviceId: deviceId,
        username: username,
        userId: userId,
        authMethod: authMethod
    }

    for (const [, sub] of subscribers.entries()) {
        sub(notify)
    }
}

if (typeof window !== "undefined") {
    token = localStorage.getItem("token")
    deviceId = localStorage.getItem("deviceId")
    username = localStorage.getItem("username")
    userId = localStorage.getItem("userId")
    authMethod = localStorage.getItem("authMethod")


    if (deviceId == null) {
        deviceId = v4()
        localStorage.setItem("deviceId", deviceId) 
    }

    console.log("token", token)
}


export const setSession = (args: {
    token: string
    username: string
    userId: string
    authMethod?: "password" | "passkey"
}) => {
    token = args.token
    username = args.username
    userId = args.userId
    authMethod = args.authMethod

    localStorage.setItem("token", token)
    localStorage.setItem("username", args.username)
    localStorage.setItem("userId", args.userId)

    if (args.authMethod) {
        localStorage.setItem("authMethod", args.authMethod)
    } else {
        localStorage.removeItem("authMethod")
    }

    notifyChanges()
}

export const clearSession = () => {
    token = undefined
    username = undefined
    userId = undefined
    authMethod = undefined

    localStorage.removeItem("token")
    localStorage.removeItem("username")
    localStorage.removeItem("userId")
    localStorage.removeItem("authMethod")

    notifyChanges()
}

export const subscribeToSessionEvents = (cb: SubscriberCallback) => {
    subscribers.add(cb)
}

export const unsubscribeToSessionEvents = (cb: SubscriberCallback) => {
    subscribers.delete(cb);
}

export const getSession = () => {
    const notify: SessionChangeNotify = {
        token: token,
        deviceId: deviceId,
        username: username,
        userId: userId,
        authMethod: authMethod
    }

    return notify
}
