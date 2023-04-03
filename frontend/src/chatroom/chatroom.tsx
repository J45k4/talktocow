import React from "react"
import { ChatroomMessages } from "./chatroom-messages"
import { ChatroomSendMessage } from "./chatroom-send-message"

import styles from "./chatroom.module.css"

export const Chatroom = (props: {
	chatroomId: string
}) => {
	return (
		<>
			<div className={styles.chatroom}>
				<ChatroomMessages chatroomId={props.chatroomId} />
			</div>
			<ChatroomSendMessage chatroomId={props.chatroomId} />
		</>
	)
}