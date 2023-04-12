import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { ApiError, getJson } from "./api-methods"
import { User, Chatroom } from "./types"

export const useGetData = <T>(path: string, def: T): [T, ApiError] => {
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

	console.log("query", router.query)

	const p: any = router.query[param] || ""

	return p
}

export const useAsync = <T>(callback: () => Promise<T>, deps: any[], def) => {
	const [value, setValue] = useState<T>(def)

	useEffect(() => {
		callback().then(setValue)
	}, deps)

	return value
}

export const useCourseMyMeta = () => {	
	const courseId = useParam("courseId")

	const v = useAsync<{
		role: number
	}>(async () => {
		const res = await getJson<any>(`/api/course/${courseId}/mymeta`)
	
		return res.payload
	}, [courseId], {})

	return v
}

export const useChatroomMembers = (chatroomId: string) => {
	const [members, setMembers] = useState<User[]>([])

	useEffect(() => {
		getJson<User[]>(`/api/chatroom/${chatroomId}/members`).then(r => {
			if (!r) {
				return
			}

			if (r.error) {
				return
			}

			if (!r.payload) {
				return
			}

			setMembers(r.payload)
		})
	}, [chatroomId, setMembers])

	return members
}

export const useUsers = () => {
	const [users, setUsers] = useState<User[]>([])

	useEffect(() => {
		getJson<User[]>(`/api/users`).then(r => {
			if (!r) {
				return
			}

			if (r.error) {
				return
			}

			if (!r.payload) {
				return
			}

			setUsers(r.payload)
		})
	}, [setUsers])

	return users
}

export const useChatrooms = () => {
	const [chatrooms, setChatrooms] = useState<Chatroom[]>([])

	useEffect(() => {
		getJson<Chatroom[]>(`/api/chatrooms`).then(r => {
			if (!r) {
				return
			}

			if (r.error) {
				return
			}

			if (!r.payload) {
				return
			}

			setChatrooms(r.payload)
		})
	}, [setChatrooms])

	return chatrooms
}

export const useChatroom = (chatroomId: string) => {
	const [chatroom, setChatroom] = useState<Chatroom>(null)

	useEffect(() => {
		getJson<Chatroom>(`/api/chatroom/${chatroomId}`).then(r => {
			if (!r) {
				return
			}

			if (r.error) {
				return
			}

			if (!r.payload) {
				return
			}

			setChatroom(r.payload)
		})
	}, [chatroomId, setChatroom])

	return chatroom
}