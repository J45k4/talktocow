import { serverUrl } from "../config";
import { getSession } from "../logic/session-manager"

export interface ServerError {
    code: number
    message: string
}

export interface ApiResponse<T> {
    statusCode: number
    error?: ServerError
    payload?: T
}

const getHeaders = () => {
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

const handleFetchResult = async <T>(r: Response): Promise<ApiResponse<T>> =>{
    const statusCode = r.status

    const jsonRes = await r.json()

    if (jsonRes.error) {
        return {
            error: jsonRes.error,
            statusCode: statusCode
        }
    }

    return {
        statusCode: statusCode,
        payload: jsonRes
    }
}

const resolveServerUrl = (path: string) => {
    if (serverUrl) {
        return serverUrl + path
    }

    return path
}

export const postJson = async <T>(path: string, payload: any): Promise<ApiResponse<T>> => {
    const headers = getHeaders()

    let res = await fetch(resolveServerUrl(path), {
        method: "POST",
        headers: headers,
        body: JSON.stringify(payload)
    })

    return handleFetchResult(res)
}

export const getJson = async <T>(path: string, query?): Promise<ApiResponse<T>> => {
    const headers = getHeaders()

    const res = await fetch(resolveServerUrl(path), {
        headers: headers,
    });

    return handleFetchResult(res)
}