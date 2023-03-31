import { useEffect } from "react";
import { useImmer } from "use-immer";
import { ChatroomMessage } from "./types";
import { chatroomMessageStore } from "./chatroom-message-store";
import { eventbus } from "./eventbus";
import { createLogger } from "./logger";

const logger = createLogger("useChatroomMessages")

export const useChatroomMessages = (chatroomId: string): ChatroomMessage[] => {
	const [messages, updateMessages] = useImmer<ChatroomMessage[]>([]);

	useEffect(() => {
		logger.info("chatroomId changed to", chatroomId);

		const initialMessages = chatroomMessageStore.getAllMessages(chatroomId);

		logger.info("initialMessages", initialMessages);

		updateMessages(initialMessages);
	}, [chatroomId, updateMessages]);

	useEffect(() => {
		const messageHandler = (msgs: ChatroomMessage[]) => {
			logger.info("messageHandler", msgs);

			updateMessages((draft) => {
				for (const msg of msgs) {
					if (msg.chatroomId === chatroomId &&
						!draft.find((m) => m.reference === msg.reference)) {

						logger.info("addMessage", msg);
						chatroomMessageStore.addMessage(msg);
						draft.unshift(msg);
					}
				}
			})
		};

		const chatroomMessagesSubject = eventbus.events("chatroomMessages");
		const subscription = chatroomMessagesSubject.subscribe(messageHandler);

		return () => {
			subscription.unsubscribe();
		};
	}, [chatroomId, updateMessages]);

	logger.info("messages", messages);

	return messages;
};
