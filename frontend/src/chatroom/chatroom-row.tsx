import { FaReply } from "react-icons/fa"
import { ChatroomMessage } from "../types"
import { formatTime } from "../utility"

import styles from "./chatroom-row.module.css"

const MessageTitle = (props: {
	userName: string
	writtenAt: string
}) => {
	return (
		<div className={styles.messageTitle}>
			<div className={styles.messageTitleLeftSide}>
				<div className={styles.messageAuthor}>
					{props.userName}
				</div>
				<div className={styles.messageTime}>
					{formatTime(props.writtenAt)}
				</div>
			</div>
			<div className={styles.messageTitleRightSide}>
				<FaReply />
			</div>
		</div>
	)
}

export const ChatroomMessageRow = (props: {
	chatroomMessage: ChatroomMessage
	grayBackground?: boolean
}) => {
	return (
		<div className={styles.messageRow} style={{
			backgroundColor: props.grayBackground ? "#f0f0f0" : "white"
		}}>
			<MessageTitle
				userName={props.chatroomMessage.userName}
				writtenAt={props.chatroomMessage.writtenAt}
			/>
			<div>
				{props.chatroomMessage.messageText}
			</div>
		</div>
	)
}
