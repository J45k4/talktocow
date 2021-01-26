import { useEffect, useState } from "react";
import { observable, Observable } from "rxjs";

export interface NewChatroomMessage {
    messageText: string
    fromUserName: string
    transmittedAt: string
}

export interface MessageFromServer {
    newChatroomMessage: NewChatroomMessage
}

export interface MessageToChatroom {
    messageText: string
    createTime: string
    transmitTime: string
}

export interface MessageToServer {
    messageToChatroom: MessageToChatroom
}

let client: WebSocket

let socketConnected = false

const callbacks = new Set();

function createClient() {
    var scheme = document.location.protocol == "https:" ? "wss" : "ws";
    var port = document.location.port ? ":" + document.location.port : "";
    var wsURL = scheme + "://" + document.location.hostname + port + "/api/socket?token=" + localStorage.getItem("token");

    client = new WebSocket(wsURL)

    client.onopen = () => {
        socketConnected = true
        console.log("Connected")
    }

    client.onclose = () => {
        socketConnected = false
        createClient();
    }

    client.onmessage = (msg) => {
        const parsedMessage = JSON.parse(msg.data)

        for (const [,cb] of callbacks.entries()) {

            (cb as any)(parsedMessage)
        }
    }
}

if (typeof document !== "undefined") {
    createClient();
}

export function subscribeToMessages(cb: (msg: MessageFromServer) => void) {
    callbacks.add(cb)
}

export function unsubscribeToMessages(cb: (msg: MessageFromServer) => void) {
    callbacks.delete(cb)
}


export const useSocketConn = () => {

    return {
        sendMessage: (message: MessageToServer) => {
            if (!socketConnected) {
                return false
            }

            client.send(JSON.stringify(message))

            return true
        }
    }
}

export const useMessages = ()  => {
    const [messages, setMessage] = useState<string[]>([])

    useEffect(() => {
        function handle(msg: MessageFromServer) {
            console.log("msg", msg)

            const transmittedAtDate = new Date(msg.newChatroomMessage.transmittedAt)

            const transmittedAtDateString = transmittedAtDate.toLocaleString()

            setMessage([
                transmittedAtDateString + " " + msg.newChatroomMessage.fromUserName + ": " + msg.newChatroomMessage.messageText,
                ...messages
            ])
        }

        subscribeToMessages(handle)

        return () => {
            unsubscribeToMessages(handle)
        }
    })

    return messages
}

export const useIsSocketConnected = () => {
    let connectd

    useEffect(() => {
        connectd = socketConnected
    }, [socketConnected])

    return connectd
}