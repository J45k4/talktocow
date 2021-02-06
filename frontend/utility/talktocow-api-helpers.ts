import { getSession } from "../logic/session-manager"

export interface ServerError {
    code: number
    message: string
}

export interface ApiResponse {
    statusCode: number
    error?: ServerError
    payload?: any
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

const handleFetchResult = async (r: Response): Promise<ApiResponse> =>{
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

export const postJson = async (path: string, payload: any): Promise<ApiResponse> => {
    const headers = getHeaders()

    let res = await fetch(path, {
        method: "POST",
        headers: headers,
        body: JSON.stringify(payload)
    })

    return handleFetchResult(res)
}

export const getJson = async (path: string, query): Promise<ApiResponse> => {
    const headers = getHeaders()

    const res = await fetch(path, {
        headers: headers,
    });

    return handleFetchResult(res)
}