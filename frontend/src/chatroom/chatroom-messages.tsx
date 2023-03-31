import React, { useEffect, useRef } from "react"
import { loadChatroomMessages } from "../chatroom-message-managers"
import { useChatroomMessages } from "../use-chatroom-messages"
import styles from "./chatroom-messages.module.css"

export const ChatroomMessages = (props: {
	chatroomId: string
}) => {
	const messages = useChatroomMessages(props.chatroomId)
	const messagesEndRef = useRef(null);

	const scrollToBottom = () => {
		messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
	};

	useEffect(() => {
		scrollToBottom();
	}, [messages]);

	useEffect(() => {
		loadChatroomMessages(props.chatroomId, 20)
	}, [props.chatroomId])

	return (
		<div className={styles.body}>
			<div ref={messagesEndRef} />
			{messages.map(p => (
				<div key={p.reference} className={styles.messageRow}>
					{p.messageText}
				</div>
			))}
		</div>
	)
}