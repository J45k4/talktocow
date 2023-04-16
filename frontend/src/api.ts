import { deleteJson, getJson, patchJson, postJson } from "./api-methods"
import { cache } from "./cache"
import { createLogger } from "./logger"
import { Chatroom, User } from "./types"
import { resolveServerUrl } from "./utility"

const logger = createLogger("api")

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
	},
	addChatroomMember: async (args: {
		chatroomId?: string
		userId?: string
	}) => {
		const res = await postJson<User>(`/api/chatroom/${args.chatroomId}/member`, {
			userId: args.userId
		})

		if (res.error) {
			return
		}

		const chatroom = cache.chatroom.get(args.chatroomId)
		
		logger.info("cache", cache, cache.getEntityMap())

		if (!chatroom) {
			return
		}

		logger.info("chatroom", chatroom)

		chatroom.members.add(res.payload)
	},
	removeChatroomMember: async (args: {
		chatroomId?: string
		userId?: string
	}) => {
		await deleteJson(`/api/chatroom/${args.chatroomId}/member/${args.userId}`)

		const chatroom = cache.chatroom.get(args.chatroomId)

		if (!chatroom) {
			return
		}

		chatroom.members.remove(args.userId)
	}
}