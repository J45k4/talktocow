import { Subject } from "rxjs";
import { serverUrl } from "../config";
import { getSession, subscribeToSessionEvents } from "./session-manager";
import { ConnEvent, MessageFromServer, MessageToServer } from "./websocket-types";

export const connEvent = new Subject<ConnEvent>()

type NewMessageSubscriberCallback = (notify: MessageFromServer) => void

export interface SocketStatusChanged {
    connected: boolean
}

type SocketStatusChangedSubscriberCallback = (notify: SocketStatusChanged) => void

const newMessageSubscribers = new Set<NewMessageSubscriberCallback>();

export const subscribeToNewMessages = (cb: NewMessageSubscriberCallback) => {
    newMessageSubscribers.add(cb)
}

export const unsubscribeToMessages = (cb: NewMessageSubscriberCallback) => {
    newMessageSubscribers.delete(cb)
}

const socketStatusChangedSubscribers = new Set<SocketStatusChangedSubscriberCallback>()

export const subscribeToSocketStatusChanged = (cb: SocketStatusChangedSubscriberCallback) => {
    socketStatusChangedSubscribers.add(cb)
}

export const unsubscribeToSocketStatusChanged = (cb: SocketStatusChangedSubscriberCallback) => {
    socketStatusChangedSubscribers.delete(cb)
}

let socket: WebSocket

let socketConnected = false
let token = getSession().token
let tryingToConnect = token != null

console.log("tryingToConnect socket " + tryingToConnect)

export const sendMessageToServer = (message: MessageToServer): boolean => {
    if (socketConnected === false) {
        return false
    }
    
    socket.send(JSON.stringify(message))   

    return true
}

const notifySocketStatusChanged = () => {
    for (const [,sub] of socketStatusChangedSubscribers.entries()) {
        sub({
            connected: socketConnected
        })
    }   
}

function createClient() {
    if (token == null) {
        return
    }

    const url = serverUrl ? new URL(serverUrl) : document.location

    var scheme = url.protocol == "https:" ? "wss" : "ws";
    var port = url.port ? ":" + url.port : "";
    var wsURL = scheme + "://" + url.hostname + port + "/api/socket?token=" + token

    console.log("wsURL", wsURL)

    socket = new WebSocket(wsURL)

    socket.onopen = () => {
        console.log("Socket onopen")
        socketConnected = true
        notifySocketStatusChanged()
    }

    socket.onclose = () => {
        console.log("socketclose")
        socketConnected = false
        notifySocketStatusChanged()

        if (tryingToConnect == false) {
            return
        }

        setTimeout(() => {
            createClient();
        }, 2000)
    }

    socket.onmessage = (msg) => {
        const parsedMessage = JSON.parse(msg.data)

        for (const [,cb] of newMessageSubscribers.entries()) {

            (cb as any)(parsedMessage)
        }

        console.log("parsedMessage", parsedMessage)

        connEvent.next({
            type: "messageFromServer",
            ...parsedMessage
        })
    }
}

subscribeToSessionEvents((s) => {
    if (s.token) {
        token = s.token
        tryingToConnect = true
        createClient()
    } else {
        console.log("Token is undefined closing connection")

        token = undefined
        tryingToConnect = false
        socketConnected = false
        socket.close(1000)

        notifySocketStatusChanged()
    }
})

if (typeof document !== "undefined") {
    createClient();
}

export const isSocketConnected = () => {
    return socketConnected
}