import { getJson, patchJson, postJson } from "./api-methods"
import { User } from "./types"
import { resolveServerUrl } from "./utility"

export const fetchChatrooms = async () => {
	return fetch(resolveServerUrl("/api/chatrooms"))
}

export const api = {
	createChatroom: async (args: {
		name?: string
		userIds?: string[]
	}) => {
		return await postJson<User>("/api/chatroom", args)
	},
	patchChatroom: async (args: {
		chatroomId?: string
		name?: string
	}) => {
		return await patchJson<User>(`/api/chatroom/${args.chatroomId}`, {
			name: args.name
		})
	},
	fetchUsers: async () => {
		await getJson("/api/users")
	}
}