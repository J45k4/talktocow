import { useCallback, useState } from "react"
import { browserSupportsWebAuthn, startRegistration } from "@simplewebauthn/browser"
import { postJson } from "./api-methods"

export type PasskeyMetadata = {
	id: number
	name: string
	authenticatorAttachment: string
	cloneWarning: boolean
	createdAt: string
	lastUsedAt?: string
}

let passkeyRegistrationInProgress = false

export const useAddPasskey = (onAdded?: () => void) => {
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState("")

	const addPasskey = useCallback(async () => {
		if (passkeyRegistrationInProgress) {
			return
		}

		if (!browserSupportsWebAuthn()) {
			setError("This browser does not support passkeys")
			return
		}

		passkeyRegistrationInProgress = true
		setLoading(true)
		setError("")

		try {
			const begin = await postJson<any>("/api/passkeys/registration/begin", {})

			if (begin.error) {
				setError(begin.error.message)
				setLoading(false)
				return
			}

			const ceremonyId = begin.payload?.ceremonyId

			if (!ceremonyId) {
				setError("Could not start passkey registration")
				return
			}

			const registration = await startRegistration({
				optionsJSON: begin.payload.options
			})

			const finish = await postJson<PasskeyMetadata>("/api/passkeys/registration/finish", {
				ceremonyId: ceremonyId,
				response: registration
			})

			if (finish.error) {
				setError(finish.error.message)
				setLoading(false)
				return
			}

			setLoading(false)
			onAdded?.()
		} catch (e) {
			setError(e instanceof Error ? e.message : "Could not add passkey")
		} finally {
			setLoading(false)
			passkeyRegistrationInProgress = false
		}
	}, [onAdded])

	return {
		addPasskey,
		loading,
		error,
		setError
	}
}
