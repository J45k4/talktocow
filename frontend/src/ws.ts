import { serverUrl } from "./config"
import { eventbus } from "./eventbus";
import { createLogger } from "./logger";
import { MessageFromServer, MessageToServer } from "./types"

const logger = createLogger("ws")

let wsSocket: WebSocket

const sendBuffer = []

const send = (msg: MessageToServer) => {
	logger.info("sendMessage", msg)

	if (wsSocket.readyState === WebSocket.OPEN) {
		wsSocket.send(JSON.stringify(msg))

		return
	}

	sendBuffer.push(msg)
}

const createConn = (token: string) => {
	const url = serverUrl ? new URL(serverUrl) : window.location

	var scheme = url.protocol == "https:" ? "wss" : "ws";
	var port = url.port ? ":" + url.port : "";
	var wsURL = `${scheme}://${url.hostname}${port}/api/ws`
	
	logger.info("wsURL", wsURL)

	wsSocket = new WebSocket(wsURL)

	wsSocket.onopen = () => {
		logger.info("Socket onopen")

		send({
			type: "authenticate",
			token: token,
			transmitedAt: new Date().toISOString()
		})

		let buffItem = sendBuffer.shift()

		while (buffItem) {
			buffItem.transmitedAt = new Date().toISOString()

			send(buffItem)

			buffItem = sendBuffer.shift()
		}
	}

	wsSocket.onclose = () => {
		logger.info("socketclose")

		setTimeout(() => {
			createConn(token)
		}, 1000)
	}

	wsSocket.onmessage = (msg) => {
		logger.info("parsedMessage", msg)

		const parsedMessage = JSON.parse(msg.data) as MessageFromServer

		if (parsedMessage.type === "chatroomMessages") {
			const messages = parsedMessage.messages

			eventbus.publish("chatroomMessages", messages)
		}
	}
}

const openConn = (token: string) => {
	if (!wsSocket) {
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

export const ws = {
	send,
	openConn
}