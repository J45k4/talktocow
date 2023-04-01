import { createLogger } from "./logger";
import { ChatroomMessage } from "./types";

const logger = createLogger("chatroomMessageStore")

const messagesByChatroom = new Map<string, ChatroomMessage[]>()

const sortMessages = (chatroomId: string) => {
	logger.info("sortMessages", chatroomId);

	const messages = messagesByChatroom.get(chatroomId);

	if (!messages) {
		return;
	}

	logger.info("messages", messages)

	// Sort messages by writtenAt so that latest messages are at the bottom
	// strings are first converted to dates and then sorted
	messages.sort((a, b) => {
		const aDate = new Date(a.writtenAt);
		const bDate = new Date(b.writtenAt);

		return aDate.getTime() - bDate.getTime();	
	})

	logger.info("sorted messages", messages)
}

export const chatroomMessageStore = {
	addMessage(message: ChatroomMessage) {
		const chatroomId = message.chatroomId;

		let chatroomMessages = messagesByChatroom.get(chatroomId);

		if (!chatroomMessages) {
			chatroomMessages = [];
			messagesByChatroom.set(chatroomId, chatroomMessages);
		}

		if (chatroomMessages.some((m) => m.reference === message.reference)) {
			return
		}

		logger.info("addMessage", message);

		chatroomMessages.push(message);

		sortMessages(chatroomId);
	},
	addMessages(messages: ChatroomMessage[]) {
		logger.info("addMessages", messages);

		const chatroomsToUpdate = new Set<string>();

		messages.forEach((message) => {
			const chatroomId = message.chatroomId;
			chatroomsToUpdate.add(chatroomId);

			let chatroomMessages = messagesByChatroom.get(chatroomId);

			if (!chatroomMessages) {
				chatroomMessages = [];
				messagesByChatroom.set(chatroomId, chatroomMessages);
			}

			if (chatroomMessages.some((m) => m.reference === message.reference)) {
				return;
			}

			chatroomMessages.push(message);
		});

		chatroomsToUpdate.forEach((chatroomId) => {
			sortMessages(chatroomId);
		});
	},
	getAllMessages(chatroomId: string): ChatroomMessage[] {
		logger.info("getAllMessages", chatroomId);

		const messages = messagesByChatroom.get(chatroomId);

		if (!messages) {
			return [];
		}

		return [...messages];
	},
};
