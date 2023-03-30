import { serverUrl } from "./config"
import { createLogger } from "./logger";
import { MessageToServer } from "./types"

const logger = createLogger("ws")

const url = serverUrl ? new URL(serverUrl) : document.location

var scheme = url.protocol == "https:" ? "wss" : "ws";
var port = url.port ? ":" + url.port : "";
var wsURL = `${scheme}://${url.hostname}${port}/api/ws`

console.log("wsURL", wsURL)

let ws: WebSocket

export const sendMessage = (msg: MessageToServer) => {
	logger.info("sendMessage", msg)

	if (ws.readyState === WebSocket.OPEN) {
		ws.send(JSON.stringify(msg))
	}
}

const createConn = (token: string) => {
	ws = new WebSocket(wsURL)

	ws.onopen = () => {
		console.log("Socket onopen")

		sendMessage({
			type: "authenticate",
			token: token
		})
	}

	ws.onclose = () => {
		console.log("socketclose")

		setTimeout(() => {
			createConn(token)
		}, 1000)
	}

	ws.onmessage = (msg) => {
		console.log("parsedMessage", msg)
	}
}

export const openConn = (token: string) => {
	if (!ws) {
		return
	}

	createConn(token)
}

if (typeof window !== "undefined") {
	const token = window.localStorage.getItem("token")

	if (token) {
		createConn(token)
	}
}