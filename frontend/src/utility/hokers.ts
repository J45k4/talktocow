import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { ApiError, getJson } from "./talktocow-api-helpers"

export const useGetData = <T>(path: string, def?: T): [T, ApiError] => {
	const [data, setData] = useState<T>(def)
	const [error, setError] = useState<ApiError>(null)

	useEffect(() => {
		getJson<T>(path).then(r => {
			if (!r) {
				return
			}

			if (r.error) {
				setError(r.error)
				return
			}

			if (!r.payload) {
				return
			}

			setData(r.payload)
		})
	}, [])

	return [data, error]
}

export const useParam = (param: string): string => {
	const router = useRouter()

	const p: any = router.query[param] || ""

	return p
}