import { useState } from "react"
import { getJson } from "./api-methods"
import { Chatroom, User } from "./types"

// type CacheMap = {
// 	"user": User
// 	"chatroom": Chatroom
// 	"chatrooms": Chatroom[]
// }

// const map = new Map<keyof CacheMap, Map<string, any>>()

// export const cache = {
// 	get: <K extends keyof CacheMap>(key: K, id: string) => {

// 	},
// 	addRelation: <K extends keyof CacheMap>(
// 		key: K, 
// 		id: string, 
// 		relationKey: string, 
// 		relationId: string
// 	) => {
	
// 	},
// 	update: <K extends keyof CacheMap>(key: K, value: CacheMap[K]) => {
// 		const m = map.get(key)

// 		if (!m) {
// 			const m = new Map<string, CacheMap[K]>()
// 			map.set(key, m)
// 		}

// 		m.set(value.id, value)
// 	},
// 	delete: <K extends keyof CacheMap>(key: K, id: string) => {
// 		const m = map.get(key)

// 		if (!m) {
// 			return
// 		}

// 		m.delete(id)
// 	},
// 	watch: <K extends keyof CacheMap>(key: K, id: string, callback: (value: CacheMap[K]) => void) => {

// 	}
// }

// const update = (id: string, u: User) => {

// }

// const watch = (query: String) => {

// }


// getJson<User[]>(`/api/chatroom/${chatroomId}/members`)
// 	.then(res => {
// 		for (const u of res.payload) {
// 			update("chatroom[0].member", u)
// 		}
// 	})

type LoadPolicy = "load-once" |
	"load-every-time"

// export type CacheField = {
// 	type: string
// 	loadPolicy?: LoadPolicy
// 	[key: string]: CacheField
// }

type TypeMap = {
	"chatroom": Chatroom
}

export const cache = {
	"chatroom": {
		"type": "chatroom",
		"load": async (root: Chatroom) => {

		},
		"create": async (root: Chatroom) => {

		},
		"messages": {
			"load": async (root: Chatroom) => {

			}
		},
		"members": {
			"type": "user",
			"load": async (root: Chatroom) => {
				const members = await getJson<User[]>(
					`/api/chatroom/${root.id}/members`)

				return members.payload
			}
		}
	},
	"chatrooms": {
		"load": async () => {
			const chatrooms = await getJson<Chatroom[]>(
				`/api/chatrooms`)

			return chatrooms.payload		
		}
	},
	"users": {
		"load": async () => {

		}
	},
	"user": {
		"load": async () => {

		}
	}
}

type t = keyof typeof cache

export const useCache = (key: string) => {
	const [value, setValue] = useState(key)

	return value
}