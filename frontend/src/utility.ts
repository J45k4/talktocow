import { serverUrl } from "./config"

export const resolveServerUrl = (path: string) => {
    if (serverUrl) {
        return serverUrl + path
    }

    return path
}