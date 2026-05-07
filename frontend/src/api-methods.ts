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
): Promise<ApiResponse<T>> => {
	const headers = getHeaders()

	try {
		const res = await fetch(resolveServerUrl(path), {
			method: method,
			headers: headers,
			body: JSON.stringify(payload)
		})

		return handleFetchResult<T>(res)
	} catch (e) {
		return {
			error: {
				code: 500,
				message: e instanceof Error ? e.message : String(e)
			}
		}
	}
}

const handleFetchResult = async <T>(r: Response): Promise<ApiResponse<T>> => {
    const statusCode = r.status

	const textRes = await r.text()
	let jsonRes: any = null

	if (textRes) {
		try {
			jsonRes = JSON.parse(textRes)
		} catch (e) {
			return {
				error: {
					code: statusCode,
					message: r.ok ? `Invalid JSON response: ${textRes}` : `Request failed with status ${statusCode}: ${textRes || r.statusText}`
				}
			}
		}
	}

	if (!jsonRes) {
		if (!r.ok) {
			return {
				error: {
					code: statusCode,
					message: `Request failed with status ${statusCode}${r.statusText ? `: ${r.statusText}` : ""}`
				}
			}
		}

		return {
			payload: null
		}
	}

    if (jsonRes.error) {
        return {
            error: jsonRes.error,
        }
    }

	if (!r.ok) {
		return {
			error: {
				code: statusCode,
				message: `Request failed with status ${statusCode}${r.statusText ? `: ${r.statusText}` : ""}`
			}
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

export const deleteJson = async <T>(path: string): Promise<ApiResponse<T>> => {
	const headers = getHeaders()

	let res = await fetch(resolveServerUrl(path), {
		method: "DELETE",
		headers: headers
	})

	return handleFetchResult(res)
}

export const getJson = async <T>(path: string, query?): Promise<ApiResponse<T>> => {
    return customFetch<T>(path, "GET")
}
