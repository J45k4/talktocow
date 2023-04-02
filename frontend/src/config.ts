import { createLogger, LogLevel, setLogLevel } from "./logger"

const logger = createLogger("config")

export const serverUrl = process.env.NEXT_PUBLIC_SERVER_URL

console.log("Serverurl ", serverUrl)

export const configLogLevel = process.env.NEXT_PUBLIC_LOG_LEVEL

if (configLogLevel === "debug") {
	setLogLevel(LogLevel.Debug)
}

if (configLogLevel === "info") {
	setLogLevel(LogLevel.Info)
}

if (configLogLevel === "warn") {
	setLogLevel(LogLevel.Warn)
}

if (configLogLevel === "error") {
	setLogLevel(LogLevel.Error)
}


export const getHTTPServerUrl = (path: string) => {
	if (serverUrl) {
		if (!path) {
			return new URL(serverUrl)
		}

		return new URL(path, serverUrl)
	}

	const href = window.location.href

	if (path) {
		return new URL(path, href)
	}

	return new URL(href)
}

export const getWebsocketServerUrl = (path: string) => {
	const serverUrl = getHTTPServerUrl(path)

	var scheme = serverUrl.protocol == "https:" ? "wss" : "ws";
	var port = serverUrl.port ? ":" + serverUrl.port : "";
	var wsURL = `${scheme}://${serverUrl.hostname}${port}`

	logger.info("wsURL", wsURL)

	return new URL(path, wsURL)
}