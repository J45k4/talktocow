import { ChatroomMessage } from "../types"
import { formatTime } from "../utility"

import styles from "./chatroom-row.module.css"

export const ChatroomMessageRow = (props: {
	chatroomMessage: ChatroomMessage
	grayBackground?: boolean
}) => {
	return (
		<div className={styles.messageRow} style={{
			backgroundColor: props.grayBackground ? "#f0f0f0" : "white"
		}}>
			<div className={styles.messageAuthor}>
				{props.chatroomMessage.userName}
			</div>
			<div className={styles.messageTime}>
				{formatTime(props.chatroomMessage.writtenAt)}
			</div>
			<div>
				{props.chatroomMessage.messageText}
			</div>
		</div>
	)
}