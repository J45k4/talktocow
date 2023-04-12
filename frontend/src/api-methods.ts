import { getSession } from "./logic/session-manager"
import { serverUrl } from "./config";
import { resolveServerUrl } from "./utility";
import { getHeaders } from "./headers";

export interface ApiError {
    code: number
    message: string
}

export interface ApiResponse<T> {
    error?: ApiError
    payload?: T
}

const customFetch = async <T>(
	path: string, 
	method: "POST" | "PUT" | "GET", 
	payload?: any
) => new Promise(async (resolve, reject) => {
	const headers = getHeaders()

	await fetch(resolveServerUrl(path), {
		method: method,
		headers: headers,
		body: JSON.stringify(payload)
	}).then(res => {
		resolve(handleFetchResult<T>(res))
	}).catch(e => {
		resolve({
			error: {
				code: 500,
				message: e.message
			}
		})
	})
})

const handleFetchResult = async <T>(r: Response): Promise<ApiResponse<T>> => {
    const statusCode = r.status

    const jsonRes = await r.json()

	if (!jsonRes) {
		return {
			payload: null
		}
	}

    if (jsonRes.error) {
        return {
            error: jsonRes.error,
        }
    }

    return {
        payload: jsonRes
    }
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

export const putJson = async <T>(path: string, payload: any): Promise<ApiResponse<T>> => {
    const headers = getHeaders()

    let res = await fetch(resolveServerUrl(path), {
        method: "PUT",
        headers: headers,
        body: JSON.stringify(payload)
    })

    return handleFetchResult(res)
}

export const patchJson = async <T>(path: string, payload: any): Promise<ApiResponse<T>> => {
	const headers = getHeaders()

	let res = await fetch(resolveServerUrl(path), {
		method: "PATCH",
		headers: headers,
		body: JSON.stringify(payload)
	})

	return handleFetchResult(res)
}

export const getJson = async <T>(path: string, query?): Promise<ApiResponse<T>> => {
    return customFetch<T>(path, "GET")
}