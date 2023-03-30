import { resolveServerUrl } from "./utility"

export const fetchChatrooms = async () => {
	return fetch(resolveServerUrl("/api/chatrooms"))
}