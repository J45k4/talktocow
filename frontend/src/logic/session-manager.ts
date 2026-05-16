import { v4 } from "uuid"
import { resolveServerUrl } from "../utility"

type SubscriberCallback = (payload: SessionChangeNotify) => void;

export interface SessionChangeNotify {
    token?: string
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
let sessionReady = Promise.resolve()

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

const clearStoredToken = () => {
    token = undefined
    localStorage.removeItem("token")
    notifyChanges()
}

const syncAuthCookieFromStoredToken = async (storedToken: string) => {
    const abortController = new AbortController()
    const timeout = window.setTimeout(() => abortController.abort(), 3000)

    try {
        const response = await fetch(resolveServerUrl("/api/session/cookie"), {
            method: "POST",
            headers: {
                Authorization: "Bearer " + storedToken
            },
            credentials: "include",
            signal: abortController.signal
        })

        if (response.ok) {
            clearStoredToken()
        }
    } catch (_error) {
    } finally {
        window.clearTimeout(timeout)
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

    if (token) {
        sessionReady = syncAuthCookieFromStoredToken(token)
    }
}


export const setSession = (args: {
    token?: string
    username: string
    userId: string
    authMethod?: "password" | "passkey"
}) => {
    token = args.token
    username = args.username
    userId = args.userId
    authMethod = args.authMethod

    if (token) {
        localStorage.setItem("token", token)
    } else {
        localStorage.removeItem("token")
    }

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

export const logout = async () => {
    try {
        await fetch(resolveServerUrl("/api/logout"), {
            method: "POST",
            credentials: "include"
        })
    } catch (e) {
        // Logout must still clear local UI state if the network request fails.
    } finally {
        clearSession()
    }
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

export const waitForSessionReady = () => sessionReady
