import React, { useEffect, useRef } from "react"
import { loadChatroomMessages } from "../chatroom-message-managers"
import { useChatroomMessages } from "../use-chatroom-messages"
import styles from "./chatroom-messages.module.css"
import { ChatroomMessageRow } from "./chatroom-row"

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
			{messages.map((p, index) => (
				<ChatroomMessageRow 
					key={p.reference} 
					chatroomMessage={p}
					grayBackground={index % 2 === 0} />
			))}
			<div ref={messagesEndRef} />
		</div>
	)
}