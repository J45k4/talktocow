import { sendMessageToServer } from "./websocket-conn"

let cooldown = false

export const sendIAmHere = () => {
    if (cooldown) {
        return
    }

    cooldown = true

    console.log("mousemove")

    sendMessageToServer({
        iamHere: true
    })

    setTimeout(() => {
        cooldown = false
    }, 10000)
}


export const startOnlineWatch = () => {
    console.log("startOnlineWatch")

    if (typeof window === "undefined") {
        return
    }

    window.onmousemove = () => {
        sendIAmHere()
    }
}