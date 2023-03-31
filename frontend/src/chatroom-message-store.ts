import { createLogger } from "./logger";
import { ChatroomMessage } from "./types";

const logger = createLogger("chatroomMessageStore")

export const chatroomMessageStore = {
	messagesByChatroom: {} as Record<number, ChatroomMessage[]>,

	addMessage(message: ChatroomMessage) {
		const chatroomId = message.chatroomId;

		let chatroomMessages = this.messagesByChatroom[chatroomId];

		if (!chatroomMessages) {
			chatroomMessages = [];
			this.messagesByChatroom[chatroomId] = chatroomMessages;
		}

		if (chatroomMessages.some((m) => m.reference === message.reference)) {
			return
		}

		logger.info("addMessage", message);

		this.messagesByChatroom[chatroomId].push(message);
		this.sortMessages(chatroomId);
	},
	addMessages(messages: ChatroomMessage[]) {
		logger.info("addMessages", messages);

		const chatroomsToUpdate = new Set<string>();

		messages.forEach((message) => {
			const chatroomId = message.chatroomId;
			chatroomsToUpdate.add(chatroomId);

			let chatroomMessages = this.messagesByChatroom[chatroomId];

			if (!chatroomMessages) {
				chatroomMessages = [];
				this.messagesByChatroom[chatroomId] = chatroomMessages;
			}

			if (chatroomMessages.some((m) => m.reference === message.reference)) {
				return;
			}

			this.messagesByChatroom[chatroomId].push(message);
		});

		chatroomsToUpdate.forEach((chatroomId) => {
			this.sortMessages(chatroomId);
		});
	},
	getAllMessages(chatroomId: string): ChatroomMessage[] {
		logger.info("getAllMessages", chatroomId);

		return this.messagesByChatroom[chatroomId] || [];
	},
	sortMessages(chatroomId: number) {
		logger.info("sortMessages", chatroomId);
	
		const messages = this.messagesByChatroom[chatroomId] || [];
		const messagesCopy = [...messages];
		messagesCopy.sort((a, b) => {
			return a.writenAt.localeCompare(b.writenAt);
		});
		this.messagesByChatroom[chatroomId] = messagesCopy;
	}
};
