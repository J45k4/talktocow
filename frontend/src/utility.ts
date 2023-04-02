import { serverUrl } from "./config"

export const resolveServerUrl = (path: string) => {
    if (serverUrl) {
        return serverUrl + path
    }

    return path
}

export const formatTimestamp = (timestamp: string) => {
	const date = new Date(timestamp)

	const year = date.getFullYear()
	const month = date.getMonth() + 1
	const day = date.getDate()
	const hours = date.getHours()
	const minutes = date.getMinutes()

	let str = `${year}-${month < 10 ? "0" + month : month}-${day < 10 ? "0" + day : day}`
	str += ` ${hours < 10 ? "0" + hours : hours}:${minutes < 10 ? "0" + minutes : minutes}`

	return str
}

export const formatTime = (timestamp: string) => {
	const date = new Date(timestamp)

	const hours = date.getHours()
	const minutes = date.getMinutes()

	return `${hours < 10 ? "0" + hours : hours}:${minutes < 10 ? "0" + minutes : minutes}`
}