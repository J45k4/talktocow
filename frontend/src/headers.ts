import { getSession } from "./logic/session-manager";

export const getHeaders = () => {
    const session = getSession();

    console.log("session", session)

    const headers = {
        ["Content-Type"]: "application/json"
    }

    if (session.token) {
        headers["Authorization"] = "Bearer " + session.token
    }

    if (session.deviceId) {
        headers["x-device-id"] = session.deviceId
    }

    return headers
}
