import { LogLevel, setLogLevel } from "./logger"

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
