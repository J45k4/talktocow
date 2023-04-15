import { getJson, patchJson, postJson } from "./api-methods"
import { cache } from "./cache"
import { chatroomState } from "./state"
import { Chatroom, User } from "./types"
import { resolveServerUrl } from "./utility"

export const fetchChatrooms = async () => {
	return fetch(resolveServerUrl("/api/chatrooms"))
}

export const api = {
	createChatroom: async (args: {
		name?: string
		userIds?: string[]
	}) => {
		return await postJson<Chatroom>("/api/chatroom", args)
	},
	patchChatroom: async (args: {
		chatroomId?: string
		name?: string
	}) => {
		const res = await patchJson<Chatroom>(`/api/chatroom/${args.chatroomId}`, {
			name: args.name
		})

		if (res.payload) {
			cache.chatroom.upsert(res.payload)
		}

		return res
	},
	fetchUsers: async () => {
		await getJson("/api/users")
	}
}