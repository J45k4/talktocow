import React from "react"
import { ChatroomMessages } from "./chatroom-messages"
import { ChatroomSendMessage } from "./chatroom-send-message"

import styles from "./chatroom.module.css"
import { BsSnapchat } from "react-icons/bs"
import { ChatroomSearchButton } from "./chatroom-search-button"
import { useChatroom, useChatroomMembers } from "../hokers"

const ChatroomTitle = (props: {
	chatroomId: string
}) => {
	const members = useChatroomMembers(props.chatroomId)
	const chatroom = useChatroom(props.chatroomId)

	return (
		<div className={styles.chatroomTitle}>
			<div className={styles.chatroomTitleLeftSide}>
				<BsSnapchat style={{
					width: "100%",
					height: "100%",
				}} />
			</div>
			<div className={styles.chatroomTitleCenter}>
				<div className={styles.chatroomChatGroup}>
					{chatroom?.name}
				</div>
				<div className={styles.chatroomChatGroupMembers}>
					{members.map(m => m.name).join(", ")}
				</div>
			</div>
			<div className={styles.chatroomTitleRightSide}>
				<ChatroomSearchButton style={{
					width: "25px",
					height: "25px",
					padding: "10px",
					paddingRight: "15px"
				}} />
			</div>
		</div>
	)
}

export const Chatroom = (props: {
	chatroomId: string
}) => {
	return (
		<div className={styles.chatroom}>
			<ChatroomTitle chatroomId={props.chatroomId} />
			<ChatroomMessages chatroomId={props.chatroomId} />
			<ChatroomSendMessage chatroomId={props.chatroomId} />
		</div>
	)
}